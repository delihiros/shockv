package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	BaseURL    string
	httpClient *http.Client
}

func New(baseURL string, port int) *Client {
	return &Client{
		BaseURL:    baseURL + ":" + strconv.Itoa(port),
		httpClient: &http.Client{},
	}
}

func (c *Client) Get(database string, key string) (string, error) {
	requestURL := path.Join(database, key)
	body, err := c.get("/"+requestURL, nil)
	if err != nil {
		return "", err
	}
	r := &GetResponse{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return "", err
	}
	if r.Status != http.StatusOK {
		return "", fmt.Errorf("not found %v/%v: status = %v", database, key, r.Status)
	}
	return r.Body, nil
}

func (c *Client) List(database string) ([]*Pair, error) {
	body, err := c.get("/"+database, nil)
	if err != nil {
		return nil, err
	}
	r := &ListResponse{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}
	if r.Status != http.StatusOK {
		return nil, fmt.Errorf("something went wrong: status = %v", r.Status)
	}
	return r.Body, nil
}

func (c *Client) Set(database string, key string, value string) error {
	body, err := c.post("/"+database, map[string]string{
		"key":   key,
		"value": value,
	})
	if err != nil {
		return err
	}
	r := &SetResponse{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return err
	}
	if r.Status != http.StatusCreated {
		return fmt.Errorf("failed to set %v/%v: status = %v", database, key, r.Status)
	}
	return nil
}

func (c *Client) Delete(database string, key string) error {
	requestURL := c.BaseURL + "/" + database + "/" + key
	_, err := c._delete(requestURL)
	return err
}

func (c *Client) NewDB(database string, diskless bool) error {
	dl := "false"
	if diskless {
		dl = "true"
	}
	body, err := c.get("/new", map[string]string{
		"name":     database,
		"diskless": dl,
	})
	r := &NewDBResponse{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return err
	}
	if r.Status != http.StatusCreated {
		return fmt.Errorf("failed to create %v: status = %v", database, r.Status)
	}
	return err
}

func (c *Client) get(endpoint string, queries map[string]string) ([]byte, error) {
	url := &url.URL{}
	q := url.Query()
	for k, v := range queries {
		q.Set(k, v)
	}
	url.RawQuery = q.Encode()
	requestURL := c.BaseURL + endpoint + url.String()
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func (c *Client) post(endpoint string, queries map[string]string) ([]byte, error) {
	values, err := json.Marshal(queries)
	if err != nil {
		return nil, err
	}
	requestURL := c.BaseURL + endpoint
	res, err := http.Post(requestURL, "application/json", bytes.NewBuffer(values))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func (c *Client) _delete(endpoint string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/delihiros/shockv/pkg/protocols"
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

func (c *Client) Get(database string, key string) (*protocols.GetResponse, error) {
	requestURL := path.Join(database, key)
	body, err := c.get("/"+requestURL, nil)
	if err != nil {
		return nil, err
	}
	r := &protocols.GetResponse{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) List(database string) (*protocols.ListResponse, error) {
	body, err := c.get("/"+database, nil)
	if err != nil {
		return nil, err
	}
	r := &protocols.ListResponse{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) Set(database string, key string, value string, ttl int) (*protocols.SetResponse, error) {
	body, err := c.post("/"+database, map[string]string{
		"key":   key,
		"value": value,
		"ttl":   strconv.Itoa(ttl),
	})
	if err != nil {
		return nil, err
	}
	r := &protocols.SetResponse{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) Delete(database string, key string) (*protocols.DeleteResponse, error) {
	requestURL := c.BaseURL + "/" + database + "/" + key
	body, err := c._delete(requestURL)
	if err != nil {
		return nil, err
	}
	r := &protocols.DeleteResponse{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) NewDB(database string, diskless bool) (*protocols.NewDBResponse, error) {
	dl := "false"
	if diskless {
		dl = "true"
	}
	body, err := c.get("/new", map[string]string{
		"database": database,
		"diskless": dl,
	})
	r := &protocols.NewDBResponse{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}
	return r, nil
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

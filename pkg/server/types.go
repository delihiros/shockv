package server

type Response struct {
	Status int `json:"status"`
}

type GetRequest struct {
	DatabaseName string `param:"dbname"`
	Key          string `param:"key"`
}

type GetResponse struct {
	*Response
	Body string `json:"body"`
}

type SetRequest struct {
	DatabaseName string `param:"dbname"`
	Key          string `form:"key" json:"key" xml:"key"`
	Value        string `form:"value" json:"value" xml:"value"`
}

type SetResponse struct {
	*Response
}

type ListRequest struct {
	DatabaseName string `param:"dbname"`
}

type ListResponse struct {
	*Response
	Body []*Pair `json:"body"`
}

type DeleteRequest struct {
	DatabaseName string `param:"dbname"`
	Key          string `param:"key"`
}

type DeleteResponse struct {
	*Response
}

type NewDBRequest struct {
	Name     string `query:"name"`
	Diskless bool   `query:"diskless"`
}

type NewDBResponse struct {
	*Response
}

type Pair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

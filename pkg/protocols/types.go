package protocols

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type GetRequest struct {
	Database string `param:"database"`
	Key      string `param:"key"`
}

type GetResponse struct {
	*Response
	Body string `json:"body"`
}

type SetRequest struct {
	Database string `param:"database"`
	Key      string `form:"key" json:"key" xml:"key"`
	Value    string `form:"value" json:"value" xml:"value"`
	TTL      string `form:"ttl" json:"ttl" xml:"ttl"`
}

type SetResponse struct {
	*Response
}

type ListRequest struct {
	Database string `param:"database"`
}

type ListResponse struct {
	*Response
	Body []*Pair `json:"body"`
}

type DeleteRequest struct {
	Database string `param:"database"`
	Key      string `param:"key"`
}

type DeleteResponse struct {
	*Response
}

type NewDBRequest struct {
	Database string `query:"database"`
	Diskless bool   `query:"diskless"`
}

type NewDBResponse struct {
	*Response
}

type Pair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

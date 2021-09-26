package client

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type GetResponse struct {
	*Response
	Body string `json:"body"`
}

type SetResponse struct {
	*Response
}

type ListResponse struct {
	*Response
	Body []*Pair `json:"body"`
}

type DeleteResponse struct {
	*Response
}

type NewDBResponse struct {
	*Response
}

type Pair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

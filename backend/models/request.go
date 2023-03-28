package models

type NewConnectionReq struct {
	Name     string `json:"name"`
	Urls     string `json:"urls"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CatIndexReq struct {
	ID string `json:"id"`
}

type GetPropertiesReq struct {
	ConnectionID string `json:"connection_id"`
	Index        string `json:"index"`
}

type QueryReq struct {
	*Pagination  `json:"pagination"`
	ConnectionID string                 `json:"connection_id"`
	Index        string                 `json:"index"`
	Condition    map[string]interface{} `json:"condition"`
}

type Pagination struct {
	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size"`
}

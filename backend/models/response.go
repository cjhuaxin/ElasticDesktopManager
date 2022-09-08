package models

type BaseResponse struct {
	ErrCode string      `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	Data    interface{} `json:"data"`
}

type ConnectionItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type IndexItem struct {
	Uuid      string `json:"uuid"`
	Index     string `json:"index"`
	Health    string `json:"health"`
	DocsCount int  `json:"docs_count"`
	StoreSize string `json:"store_size"`
}

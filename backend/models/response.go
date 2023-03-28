package models

import "encoding/json"

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
	DocsCount int64  `json:"docs_count"`
	StoreSize string `json:"store_size"`
}

type SearchRes struct {
	Records json.RawMessage `json:"records"`
}

type PropertieItem struct {
	Title   string `json:"title"`
	Key     string `json:"key"`
	DataKey string `json:"dataKey"`
}

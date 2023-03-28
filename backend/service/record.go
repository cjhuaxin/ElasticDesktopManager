package service

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cjhuaxin/ElasticDesktopManager/backend/base"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/errcode"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/models"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/util"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/tidwall/gjson"
)

type Record struct {
	*base.Service
}

func NewRecord(baseService *base.Service) *Record {
	return &Record{
		Service: baseService,
	}
}

func (r *Record) Init(ctx *models.EdmContext) error {
	r.Ctx = ctx

	return nil
}

func (r *Record) GetProperties(req *models.GetPropertiesReq) *models.BaseResponse {
	client, err := GetConnectionById(r.Service, req.ConnectionID)
	if err != nil {
		r.Log.Errorf("get connection[%s] failed:%v", req.ConnectionID, err)
		return r.BuildFailed(errcode.DatabaseErr, err.Error())
	}
	properties, err := r.assembleProperties(req, client)
	if err != nil {
		r.Log.Errorf("GetProperties[%s] failed:%v", req.ConnectionID, err)
		return r.BuildFailed(errcode.DatabaseErr, err.Error())
	}
	propItemList := make([]*models.PropertieItem, 0, len(properties))
	for _, p := range properties {
		propItemList = append(propItemList, &models.PropertieItem{
			Title:   p,
			Key:     p,
			DataKey: p,
		})
	}

	return r.BuildSucess(propItemList)
}

func (r *Record) Search(req *models.QueryReq) *models.BaseResponse {
	fmt.Printf("requests: %#v\n", *req)
	client, err := GetConnectionById(r.Service, req.ConnectionID)
	if err != nil {
		r.Log.Errorf("get connection[%s] failed:%v", req.ConnectionID, err)
		return r.BuildFailed(errcode.DatabaseErr, err.Error())
	}
	var esResp *esapi.Response
	if len(req.Condition) == 0 {
		esResp, err = r.queryAllRecord(client, req.Index)
		if err != nil {
			r.Log.Errorf("queryAllRecord[%s] failed:%v", req.Index, err)
			return r.BuildFailed(errcode.DatabaseErr, err.Error())
		}
	}
	body := util.ReadEsBody(esResp.Body)

	resp := &models.SearchRes{
		Records: json.RawMessage(body),
	}

	return r.BuildSucess(resp)
}

func (r *Record) assembleConditions(req *models.QueryReq) map[string]interface{} {
	if len(req.Condition) == 0 {
		return nil
	}

	return nil
}

func (r *Record) assembleProperties(req *models.GetPropertiesReq, client *elasticsearch.Client) ([]string, error) {
	resp, err := client.Indices.GetMapping(client.Indices.GetMapping.WithIndex(req.Index))
	if err != nil {
		return nil, err
	}

	body := util.ReadEsBody(resp.Body)
	properties := gjson.Get(body, fmt.Sprintf("%s.mappings.properties", req.Index)).Map()
	columns := make([]string, 0, len(properties))
	for k := range properties {
		if !strings.HasPrefix(k, "@") {
			columns = append(columns, k)
		}
	}

	return columns, nil
}

func (r *Record) queryAllRecord(client *elasticsearch.Client, index string) (*esapi.Response, error) {
	body := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}
	return client.Search(
		client.Search.WithIndex(index),
		client.Search.WithBody(esutil.NewJSONReader(&body)),
	)
}

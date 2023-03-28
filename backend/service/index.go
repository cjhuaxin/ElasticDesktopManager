package service

import (
	"fmt"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/base"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/errcode"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/models"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/resource"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/util"
	"github.com/tidwall/gjson"
)

type Index struct {
	*base.Service
}

func NewIndex(baseService *base.Service) *Index {
	return &Index{
		Service: baseService,
	}
}

func (i *Index) Init(ctx *models.EdmContext) error {
	i.Ctx = ctx

	return nil
}

func (i *Index) CatIndex(req *models.CatIndexReq) *models.BaseResponse {
	client, err := GetConnectionById(i.Service, req.ID)
	if err != nil {
		i.Log.Errorf("get connection[%s] failed:%v", req.ID, err)
		return i.BuildFailed(errcode.DatabaseErr, err.Error())
	}
	response, err := client.Cat.Indices(client.Cat.Indices.WithFormat(resource.EsDefaultFormat))
	if err != nil {
		i.Log.Errorf("cat indeices failed:%v", err)
		return i.BuildFailed(errcode.DatabaseErr, err.Error())
	}

	body := util.ReadEsBody(response.Body)
	shownIndices := make([]*models.IndexItem, 0)

	var index int
	for {
		result := gjson.Get(body, fmt.Sprint(index))
		if result.Type == gjson.Null {
			break
		}

		shownIndices = append(shownIndices, &models.IndexItem{
			Uuid:      result.Get("uuid").Str,
			Index:     result.Get("index").Str,
			Health:    result.Get("health").Str,
			DocsCount: result.Get("docs\\.count").Int(),
			StoreSize: result.Get("store\\.size").Str,
		})
		index++
	}

	return i.BuildSucess(shownIndices)
}

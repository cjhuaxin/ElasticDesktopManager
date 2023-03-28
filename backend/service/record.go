package service

import (
	"github.com/cjhuaxin/ElasticDesktopManager/backend/base"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/models"
	"github.com/olivere/elastic/v7"
)

type Record struct {
	*base.BaseService
}

func NewRecord(baseService *base.BaseService) *Record {
	return &Record{
		BaseService: baseService,
	}
}

func (r *Record) Init(ctx *models.EdmContext) error {
	r.Ctx = ctx

	return nil
}

func (r *Record) Query(req *models.QueryReq) *models.BaseResponse {
	var query elastic.Query
	if req ==nil{
		query = elastic.NewMatchAllQuery()
	}

	r.Ctx.GetEsClient().
}

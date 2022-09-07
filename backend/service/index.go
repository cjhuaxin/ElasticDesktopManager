package service

import (
	"github.com/cjhuaxin/ElasticDesktopManager/backend/base"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/models"
)

type Index struct {
	*base.BaseService
}

func NewIndex(baseService *base.BaseService) *Index {
	return &Index{
		BaseService: baseService,
	}
}

func (i *Index) Init(ctx *models.EdmContext) error {
	i.Ctx = ctx

	return nil
}

func (i *Index) CatIndex(req *models.CatIndexReq) *models.BaseResponse {
	client := i.Ctx.GetEsClient(req.ID)
	if client ==nil{
		
		i.InitEsClient()
	}
}

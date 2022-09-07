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

func (c *Index) Init(ctx *models.EdmContext) error {
	c.Ctx = ctx
	

	return nil
}
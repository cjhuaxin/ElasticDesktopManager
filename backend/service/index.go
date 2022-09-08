package service

import (
	"context"
	"strings"

	"github.com/cjhuaxin/ElasticDesktopManager/backend/base"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/errcode"
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
	if client == nil {
		stmt, err := i.Ctx.GetDbClient().Prepare("SELECT id,name,urls,user from connection WHERE id = ?")
		if err != nil {
			i.Log.Errorf("prepare select connection sql failed:%v", err)
			return i.BuildFailed(errcode.DatabaseErr, err.Error())
		}
		defer stmt.Close()
		var id, name, urls, user string
		err = stmt.QueryRow(req.ID).Scan(&id, &name, &urls, &user)
		if err != nil {
			i.Log.Errorf("query connection failed:%v", err)
			return i.BuildFailed(errcode.DatabaseErr, err.Error())
		}
		item, err := i.Keyring.Get(req.ID)
		if err != nil {
			i.Log.Errorf("get encrypt aes key from keyring failed:%v", err)
			return i.BuildFailed(errcode.DatabaseErr, err.Error())
		}
		client, err = i.InitEsClient(urls, user, string(item.Data))
		if err != nil {
			i.Log.Errorf("init es client failed:%v", err)
			return i.BuildFailed(errcode.DatabaseErr, err.Error())
		}

		i.Ctx.SetEsClient(req.ID, client)
	}

	response, err := client.CatIndices().Do(context.TODO())
	if err != nil {
		i.Log.Errorf("cat indeices failed:%v", err)
		return i.BuildFailed(errcode.DatabaseErr, err.Error())
	}
	shownIndices := make([]*models.IndexItem, 0)
	for _, index := range response {
		if strings.HasPrefix(index.Index, ".") {
			continue
		}
		shownIndices = append(shownIndices, &models.IndexItem{
			Uuid:      index.UUID,
			Index:     index.Index,
			Health:    index.Health,
			DocsCount: index.DocsCount,
			StoreSize: strings.ToUpper(index.StoreSize),
		})
	}

	return i.BuildSucess(shownIndices)
}

package service

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/cjhuaxin/ElasticDesktopManager/backend/base"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/errcode"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/models"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/resource"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/util"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/xid"
)

type Connection struct {
	*base.BaseService
}

func NewConnection(baseService *base.BaseService) *Connection {
	return &Connection{
		BaseService: baseService,
	}
}

func (c *Connection) Init(ctx *models.EdmContext) error {
	c.Ctx = ctx
	err := c.initDbClient()
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) TestEsConnection(req *models.NewConnectionReq) *models.BaseResponse {
	_, err := c.InitEsClient(req.Urls, req.Username, req.Password)
	if err != nil {
		return c.BuildFailed(errcode.ConnectionErr, err.Error())
	}

	return c.BuildSucess(nil)
}

func (c *Connection) CreateEsConnection(req *models.NewConnectionReq) *models.BaseResponse {
	client, err := c.InitEsClient(req.Urls, req.Username, req.Password)
	if err != nil {
		c.Log.Errorf("init es client failed: %v", err)
		return c.BuildFailed(errcode.ConnectionErr, err.Error())
	}
	id := xid.New().String()
	aesKey, err := c.Keyring.Get(resource.EncryptAESKey)
	if err != nil {
		c.Log.Errorf("get aes key failed: %v", err)
		return c.BuildFailed(errcode.KeyringErr, err.Error())
	}
	encryptedPwd, err := util.AesEncrypt(aesKey.Data, req.Password)
	if err != nil {
		c.Log.Errorf("aes encrypt failed: %v", err)
		return c.BuildFailed(errcode.CommonErr, err.Error())
	}
	urls, err := util.NormalizeUrls(req.Urls)
	if err != nil {
		c.Log.Errorf("normalize urls failed: %v", err)
		return c.BuildFailed(errcode.CommonErr, err.Error())
	}
	_, err = c.dbClient.Exec(
		fmt.Sprintf("INSERT INTO connection(id,name,urls,user,password) values('%s','%s','%s','%s','%s')",
			id, req.Name, strings.Join(urls, ","), req.Username, encryptedPwd))
	if err != nil {
		c.Log.Errorf("save the connection info to db failed: %v", err)
		return c.BuildFailed(errcode.DatabaseErr, err.Error())
	}
	c.Ctx.SetEsClient(id, client)

	return c.BuildSucess(nil)
}

func (c *Connection) GetSavedConnectionList() *models.BaseResponse {
	rows, err := c.dbClient.Query("SELECT id, name from connection")
	if err != nil {
		c.Log.Errorf("query connection failed: %v", err)
		return c.BuildFailed(errcode.DatabaseErr, err.Error())
	}
	defer rows.Close()
	connectionList := make([]*models.ConnectionList, 0)
	for rows.Next() {
		var id string
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			c.Log.Errorf("query connection failed: %v", err)
			continue
		}
		connectionList = append(connectionList, &models.ConnectionList{
			ID:   id,
			Name: name,
		})
	}

	return c.BuildSucess(connectionList)
}

func (c *Connection) initDbClient() error {
	db, err := sql.Open("sqlite3", filepath.Join(c.Paths.DbDir, "edm.db"))
	if err != nil {
		return err
	}
	_, err = db.Exec(resource.CreateConnectionTableSql)
	if err != nil {
		return err
	}
	c.dbClient = db

	return nil
}

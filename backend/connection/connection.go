package connection

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/cjhuaxin/ElasticDesktopManager/backend/models"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/resource"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/util"
	"github.com/goware/urlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/olivere/elastic/v7"
	"github.com/rs/xid"
)

type Connection struct {
	ctx         *models.EdmContext
	esClientMap map[string]*elastic.Client
	dbClient    *sql.DB
}

func NewConnection() *Connection {
	return &Connection{}
}

func (c *Connection) Init(ctx *models.EdmContext) error {
	c.ctx = ctx
	err := c.initDbClient()
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) TestEsConnection(req *models.NewConnectionReq) string {
	_, err := c.initEsClient(req)
	if err != nil {
		return err.Error()
	}

	return ""
}

func (c *Connection) CreateEsConnection(req *models.NewConnectionReq) string {
	client, err := c.initEsClient(req)
	if err != nil {
		c.ctx.Log.Errorf("init es client failed: %v", err)
		return err.Error()
	}
	id := xid.New().String()
	aesKey, err := c.ctx.Keyring.Get(resource.EncryptAESKey)
	if err != nil {
		c.ctx.Log.Errorf("get aes key failed: %v", err)
		return err.Error()
	}
	encryptedPwd, err := util.AesEncrypt(aesKey.Data, req.Password)
	if err != nil {
		c.ctx.Log.Errorf("aes encrypt failed: %v", err)
		return err.Error()
	}
	urls, err := normalizeUrls(req.Urls)
	if err != nil {
		c.ctx.Log.Errorf("normalize urls failed: %v", err)
		return err.Error()
	}
	_, err = c.dbClient.Exec(
		fmt.Sprintf("INSERT INTO connection(id,name,urls,user,password) values('%s','%s','%s','%s','%s')",
			id, req.Name, strings.Join(urls, ","), req.Username, encryptedPwd))
	if err != nil {
		c.ctx.Log.Errorf("save the connection info to db failed: %v", err)
		return err.Error()
	}
	if c.esClientMap == nil {
		c.esClientMap = make(map[string]*elastic.Client)
	}
	c.esClientMap[id] = client

	return ""
}

func (c *Connection) GetEsClient(connectionId string) *elastic.Client {
	return c.esClientMap[connectionId]
}

func (c *Connection) initEsClient(req *models.NewConnectionReq) (*elastic.Client, error) {
	options := make([]elastic.ClientOptionFunc, 0)
	urls, err := normalizeUrls(req.Urls)
	if err != nil {
		return nil, err
	}
	options = append(options, elastic.SetURL(urls...))
	if req.Username != "" {
		options = append(options, elastic.SetBasicAuth(req.Username, req.Password))
	}
	options = append(options, elastic.SetSniff(false))

	return elastic.NewClient(options...)
}

func (c *Connection) initDbClient() error {
	db, err := sql.Open("sqlite3", filepath.Join(c.ctx.Paths.DbDir, "edm.db"))
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

func normalizeUrls(urls string) ([]string, error) {
	urlSlice := strings.Split(urls, ",")
	endpoints := make([]string, 0)
	for _, addr := range urlSlice {
		parsed, err := urlx.Parse(addr)
		if err != nil {
			return nil, err
		}
		host, port, err := urlx.SplitHostPort(parsed)
		if err != nil {
			return nil, err
		}
		if port == "" {
			port = resource.EsDefaultPort
		}
		endpoints = append(endpoints, fmt.Sprintf("%s://%s:%s", parsed.Scheme, host, port))
	}

	return endpoints, nil
}

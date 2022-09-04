package backend

import (
	"context"
	"database/sql"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/dto"
	"github.com/olivere/elastic/v7"
	"path/filepath"
	"strings"
)

type Connection struct {
	ctx         context.Context
	esClientMap map[string]*elastic.Client
}

func NewConnection() *Connection {
	return &Connection{}
}

func (c *Connection) Init(ctx context.Context) error {
	c.ctx = ctx

	return nil
}

func (*Connection) TestEsConnection(req *dto.NewConnectionReq) string {
	_, err := initEsClient(req)
	if err != nil {
		return err.Error()
	}

	return ""
}

func (c *Connection) CreateEsConnection(req *dto.NewConnectionReq) string {
	client, err := initEsClient(req)
	if err != nil {
		return err.Error()
	}
	c.esClient = client

	return ""
}

func (c *Connection) GetEsClient() *elastic.Client {
	return c.esClient
}

func (c *Connection) initEsClient(req *dto.NewConnectionReq) (*elastic.Client, error) {
	options := make([]elastic.ClientOptionFunc, 0)
	adders := strings.Split(req.Urls, ",")
	options = append(options, elastic.SetURL(adders...))
	if req.Username != "" {
		options = append(options, elastic.SetBasicAuth(req.Username, req.Password))
	}
	options = append(options, elastic.SetSniff(false))
	return elastic.NewClient(options...)
}

func (c *Connection) initDbClient() {
	db, err := sql.Open("sqlite3", filepath.Join(c.ctx.GetPath().DbDir, ""))
	if err != nil {

	}
	defer db.Close()
}

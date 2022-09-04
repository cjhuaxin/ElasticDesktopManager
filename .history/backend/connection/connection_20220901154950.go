package connection

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/cjhuaxin/ElasticDesktopManager/backend/models"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/resource"
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
		return err.Error()
	}
	id := xid.New().String()
	_, err = c.dbClient.Exec(fmt.Sprintf("INSERT INTO connection(id,name,urls,user,password) values(%s,%s,%s,%s,%s)",
		id, req.Name, req.Urls, req.Username, req.Username))
	if err != nil {
		return err.Error()
	}
	c.esClientMap[id] = client

	return ""
}

func (c *Connection) GetEsClient(connectionId string) *elastic.Client {
	return c.esClientMap[connectionId]
}

func (c *Connection) initEsClient(req *models.NewConnectionReq) (*elastic.Client, error) {
	options := make([]elastic.ClientOptionFunc, 0)
	adders := strings.Split(req.Urls, ",")
	options = append(options, elastic.SetURL(adders...))
	if req.Username != "" {
		options = append(options, elastic.SetBasicAuth(req.Username, req.Password))
	}
	options = append(options, elastic.SetSniff(false))
	return elastic.NewClient(options...)
}

func (c *Connection) initDbClient() error {
	db, err := sql.Open("sqlite3", filepath.Join(c.ctx.GetPath().DbDir, "edm.db"))
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

package models

import (
	"database/sql"
	"time"

	"github.com/elastic/go-elasticsearch/v8"

	"golang.org/x/net/context"
)

type EdmContext struct {
	ctxMap      map[string]interface{}
	esClientMap map[string]*elasticsearch.Client
	dbClient    *sql.DB
	originCtx   context.Context
}

func (*EdmContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*EdmContext) Done() <-chan struct{} {
	return nil
}

func (*EdmContext) Err() error {
	return nil
}

func NewContext(ctx context.Context) *EdmContext {
	return &EdmContext{
		originCtx: ctx,
	}
}

func (c *EdmContext) Value(key interface{}) interface{} {
	if k, ok := key.(string); ok {
		return c.ctxMap[k]
	}

	return nil
}

func (c *EdmContext) SetValue(key string, value interface{}) {
	c.ctxMap[key] = value
}

func (c *EdmContext) GetOriginCtx() context.Context {
	return c.originCtx
}

func (c *EdmContext) GetEsClient(id string) *elasticsearch.Client {
	if c.esClientMap == nil || len(c.esClientMap) == 0 {
		return nil
	}

	return c.esClientMap[id]
}

func (c *EdmContext) SetEsClient(id string, client *elasticsearch.Client) {
	if c.esClientMap == nil {
		c.esClientMap = make(map[string]*elasticsearch.Client)
	}

	c.esClientMap[id] = client
}

func (c *EdmContext) GetDbClient() *sql.DB {
	return c.dbClient
}

func (c *EdmContext) SetDbClient(client *sql.DB) {
	c.dbClient = client
}

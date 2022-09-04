package backend

import (
	"github.com/cjhuaxin/ElasticDesktopManager/backend/connection"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"time"
)

type EdmContext struct {
	conn      *connection.Connection
	ctxMap    map[string]interface{}
	originCtx context.Context
	paths     *paths
	log       *zap.SugaredLogger
}

type paths struct {
	homeDir string
	confDir string
	dbDir   string
	logDir  string
	tmpDir  string
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

func (c *EdmContext) Value(key interface{}) interface{} {
	if k, ok := key.(string); ok {
		return c.ctxMap[k]
	}

	return nil
}

func NewContext(ctx context.Context) *EdmContext {
	return &EdmContext{
		originCtx: ctx,
	}
}

func (c *EdmContext) SetValue(key string, value interface{}) {
	c.ctxMap[key] = value
}

func (c *EdmContext) GetConn() *connection.Connection {
	return c.conn
}

func (c *EdmContext) SetConn(conn *connection.Connection) {
	c.conn = conn
}

func (c *EdmContext) GetOriginCtx() context.Context {
	return c.originCtx
}

func (c *EdmContext) GetPath() *paths {
	return c.paths
}

func (c *EdmContext) Log() *zap.SugaredLogger {
	return c.log
}

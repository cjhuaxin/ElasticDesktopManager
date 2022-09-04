package models

import (
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"time"
)

type EdmContext struct {
	ctxMap    map[string]interface{}
	originCtx context.Context
	paths     *Paths
	Log       *zap.SugaredLogger
	
}

type Paths struct {
	HomeDir string
	ConfDir string
	DbDir   string
	LogDir  string
	TmpDir  string
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

func (c *EdmContext) GetOriginCtx() context.Context {
	return c.originCtx
}

func (c *EdmContext) GetPath() *Paths {
	return c.paths
}

func (c *EdmContext) SetPath(paths *Paths) {
	c.paths = paths
}

func (c *EdmContext) Log() *zap.SugaredLogger {
	return c.log
}

func (c *EdmContext) SetLog(log *zap.SugaredLogger) {
	c.log = log
}

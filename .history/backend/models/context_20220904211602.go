package models

import (
	"time"

	"github.com/99designs/keyring"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

type EdmContext struct {
	ctxMap    map[string]interface{}
	originCtx context.Context
	Paths     *Paths
	Log       *zap.SugaredLogger
	Keyring   *keyring.Keyring
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

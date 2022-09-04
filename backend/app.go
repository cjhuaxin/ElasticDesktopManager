package main

import (
	"context"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/dto"
)

// App struct
type App struct {
	ctx *dto.EdmContext
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) onStart(ctx context.Context) {
	c := dto.NewContext(ctx)
	a.ctx = c
}

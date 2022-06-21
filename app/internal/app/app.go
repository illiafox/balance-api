package app

import (
	"balance-service/app/internal/config"
	"balance-service/app/pkg/closer"
	"balance-service/app/pkg/logger"
)

type App struct {
	flags flags
	//
	logger logger.Logger
	cfg    config.Config
	//
	closers closer.Closers
}

func (app *App) Run() {
	app.ReadConfig()
	app.Listen()
}

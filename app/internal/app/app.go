package app

import (
	"balance-service/app/internal/config"
	"go.uber.org/zap"
)

type App struct {
	flags flags
	//
	logger *zap.Logger
	cfg    config.Config
	//
	closers struct {
		logger, db func() error
	}
}

func (app *App) Run() {
	app.ReadConfig()
	app.Listen()
}

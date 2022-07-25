package app

import (
	"os"

	"balance-service/app/internal/config"
	"balance-service/app/pkg/logger"
)

func (app *App) ReadConfig() {
	// // config
	cfg, err := config.New()
	if err != nil {
		app.logger.Error("read config", logger.Error(err))

		// close logger
		app.closers.Close(app.logger)

		os.Exit(1)
	}

	app.cfg = cfg
}

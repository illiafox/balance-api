package app

import (
	"os"

	"balance-service/app/internal/config"
	"go.uber.org/zap"
)

func (app *App) ReadConfig() {
	// // config
	cfg, err := config.New()
	if err != nil {
		app.logger.Error("read config", zap.Error(err))

		// close logger
		app.closers.Close()

		os.Exit(1)
	}

	app.cfg = cfg
}

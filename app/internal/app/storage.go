package app

import (
	"os"
	"time"

	"go.uber.org/zap"
	links2 "ozon-url-shortener/app/internal/adapters/db/links"
	"ozon-url-shortener/app/internal/domain/links"
	"ozon-url-shortener/app/pkg/client/redis"
)

func (app *App) Storage() links.Storage {
	app.logger.Info("Initializing storage")

	var storage links.Storage

	if app.flags.cache {
		app.logger.Warn("Using built-in storage")
		// cache
		storage = links2.NewMemStorage()
	} else {
		client, err := redis.New(redis.Config(app.cfg.Redis), time.Second*5)
		if err != nil {
			app.logger.Error("init redis",
				zap.Error(err),
			)

			// close logger
			err = app.closers.logger()
			if err != nil {
				app.logger.Error("close logger",
					zap.Error(err),
				)
			}

			os.Exit(1)
		}

		storage = links2.NewRedisStorage(client)
		app.closers.db = client.Close
	}

	return storage
}

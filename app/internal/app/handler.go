package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"balance-service/app/internal/composites"
	v1 "balance-service/app/internal/controller/http/v1"
	_ "balance-service/docs"
	swagger "github.com/swaggo/http-swagger"
)

func (app *App) Handler() (http.Handler, error) {
	app.logger.Info("Initializing storages")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	pg, err := composites.NewPgComposite(ctx, app.cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", err)
	}
	app.closers.Add(pg)

	r, err := composites.NewRedisComposite(ctx, app.cfg.Redis)
	if err != nil {
		return nil, fmt.Errorf("redis: %w", err)
	}
	app.closers.Add(r)
	// //

	app.logger.Info("Initializing handlers")

	// //
	balance, err := composites.NewBalanceComposite(pg, r)
	if err != nil {
		return nil, fmt.Errorf("create balance composite: %w", err)
	}
	// //

	// // Routing

	router := http.NewServeMux()
	{
		api := http.NewServeMux()
		api.Handle("/v1/", http.StripPrefix("/v1", v1.New(app.logger.Named("api/v1"), balance)))
		//
		router.Handle("/api/", http.StripPrefix("/api", api))
	}

	if !app.flags.noswag {
		// swagger
		router.Handle("/swagger/", swagger.Handler())
	}

	return router, nil
}

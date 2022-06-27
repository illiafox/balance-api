package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"balance-service/app/internal/composites"
	_ "balance-service/docs"
	swagger "github.com/swaggo/http-swagger"
)

func (app *App) Handler() (http.Handler, error) {
	app.logger.Info("Initializing storages")

	// //
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
	v1, err := composites.NewBalanceComposite(app.logger.Named("api/v1"), pg, r)
	if err != nil {
		return nil, fmt.Errorf("create balance composite: %w", err)
	}
	// //

	// // Routing
	root := http.NewServeMux()
	// api/v1
	root.Handle("/api/v1/", http.StripPrefix("/api/v1", v1.Handler))

	if !app.flags.noswag {
		// swagger
		root.Handle("/swagger/", swagger.Handler())
	}
	// //

	return root, nil
}

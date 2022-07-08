package app

import (
	"context"
	"fmt"
	"net/http"
	"net/http/pprof"
	"time"

	"balance-service/app/internal/composites"
	api "balance-service/app/internal/controller/http"
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

	// API
	router.Handle("/api/", http.StripPrefix("/api",
		api.New(app.logger.Named("api/handlers"), balance)),
	)

	// Swagger
	if app.flags.swagger { // swagger
		router.Handle("/swagger/", swagger.Handler())
	}

	// pprof
	if app.flags.pprof { // pprof
		router.HandleFunc("/debug/pprof/", pprof.Index)
		router.HandleFunc("/debug/pprof/heap", pprof.Index)
		router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		router.HandleFunc("/debug/pprof/profile", pprof.Profile)
		router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		router.HandleFunc("/debug/pprof/trace", pprof.Trace)
	}

	return router, nil
}

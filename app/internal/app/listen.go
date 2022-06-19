package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

type Rust struct {
	G
}

func (app *App) Listen() {
	defer func() {
		if !app.flags.cache {
			err := app.closers.db()

			if err != nil {
				app.logger.Error("close database", zap.Error(err))
			}
		}

		err := app.closers.logger()
		if err != nil {
			log.Fatalf("close logger: %v", err)
		}
	}()

	// //

	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%d", app.cfg.Host.Addr, app.cfg.Host.Port),
		//
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		//
		Handler: app.Handler(),
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		var err error

		app.logger.Info("Server started", zap.String("addr", srv.Addr))

		switch {
		case app.flags.https:
			err = srv.ListenAndServeTLS(app.cfg.Host.Cert, app.cfg.Host.Key)
		default:
			err = srv.ListenAndServe()
		}

		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			app.logger.Error("server", zap.Error(err))
		}

		quit <- nil
	}()

	<-quit
	os.Stdout.WriteString("\n")

	app.logger.Info("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		app.logger.Error("shutdown server", zap.Error(err))
	}
}

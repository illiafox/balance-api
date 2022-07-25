package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"balance-service/app/pkg/logger"
)

func (app *App) Listen() {
	handler, err := app.Handler()
	if err != nil {
		app.logger.Error("create handler", logger.Error(err))

		app.closers.Close(app.logger)
		os.Exit(1)
	}
	defer app.closers.Close(app.logger)

	// //

	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%d", app.cfg.Host.Addr, app.cfg.Host.Port),
		//
		Handler: handler,
	}

	if !app.flags.pprof {
		srv.WriteTimeout = time.Second * 3
		srv.ReadTimeout = time.Second * 3
		srv.IdleTimeout = time.Second * 15
	}

	// //
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	// //

	go func() {
		app.logger.Info("Server started",
			logger.String("address", srv.Addr),
			logger.Bool("https", app.flags.https),
		)

		var err error
		switch {
		case app.flags.https:
			err = srv.ListenAndServeTLS(app.cfg.Host.Cert, app.cfg.Host.Key)
		default:
			err = srv.ListenAndServe()
		}

		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			app.logger.Error("server", logger.Error(err))
		}

		stop()
	}()

	// //

	<-ctx.Done() // wait for signal or nil
	_, _ = os.Stdout.WriteString("\n")

	// //

	app.logger.Info("Shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		app.logger.Error("shutdown", logger.Error(err))
	}
}

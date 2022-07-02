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

	"go.uber.org/zap"
)

func (app *App) Listen() {
	handler, err := app.Handler()
	if err != nil {
		app.logger.Error("create handler", zap.Error(err))

		app.closers.Close()
		os.Exit(1)
	}
	defer app.closers.Close()

	// //

	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%d", app.cfg.Host.Addr, app.cfg.Host.Port),
		//
		WriteTimeout: time.Second * 3,
		ReadTimeout:  time.Second * 3,
		IdleTimeout:  time.Second * 15,
		//
		Handler: handler,
	}

	// //
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	// //

	go func() {
		app.logger.Info("Server started",
			zap.String("address", srv.Addr),
			zap.Bool("https", app.flags.https),
		)

		var err error
		switch {
		case app.flags.https:
			err = srv.ListenAndServeTLS(app.cfg.Host.Cert, app.cfg.Host.Key)
		default:
			err = srv.ListenAndServe()
		}

		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			app.logger.Error("server", zap.Error(err))
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
		app.logger.Error("shutdown", zap.Error(err))
	}
}

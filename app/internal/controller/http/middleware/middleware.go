package middleware

import (
	"context"
	"net/http"
	"time"

	"balance-service/app/pkg/logger"
	"go.uber.org/zap"
)

type Middleware struct {
	l logger.Logger
	t time.Duration
}

func (m Middleware) Use(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// logger
		ctx := WithLogger(r.Context(), m.l.With(zap.String("endpoint", r.URL.Path)))
		// timeout
		ctx, cancel := context.WithTimeout(ctx, m.t)
		defer cancel()
		//
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

func New(logger logger.Logger, timeout time.Duration) Middleware {
	return Middleware{
		l: logger,
		t: timeout,
	}
}

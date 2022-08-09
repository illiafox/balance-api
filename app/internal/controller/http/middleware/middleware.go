package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"balance-service/app/pkg/logger"
)

type Middleware struct {
	l logger.Logger
	t time.Duration
}

func (m Middleware) Use(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// logger
		ctx := WithLogger(r.Context(), m.l.With(logger.String("endpoint", r.URL.Path)))
		// timeout
		ctx, cancel := context.WithTimeout(ctx, m.t)
		defer cancel()
		//
		t := time.Now()
		handler.ServeHTTP(w, r.WithContext(ctx))
		fmt.Println(time.Since(t).String())
	})
}

func New(logger logger.Logger, timeout time.Duration) Middleware {
	return Middleware{
		l: logger,
		t: timeout,
	}
}

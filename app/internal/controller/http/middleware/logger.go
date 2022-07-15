package middleware

import (
	"context"

	"balance-service/app/pkg/logger"
)

type loggerKey struct{}

func WithLogger(ctx context.Context, logger logger.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func GetLogger(ctx context.Context) logger.Logger {
	l, _ := ctx.Value(loggerKey{}).(logger.Logger)
	return l
}

package cache

import (
	"context"
	"errors"

	"github.com/shopspring/decimal"
)

var ErrBalanceNotFound = errors.New("balance not found")

type CacheStorage interface {
	GetCachedBalance(ctx context.Context, userID int64) (decimal decimal.Decimal, err error)
	UpdateCachedBalance(ctx context.Context, userID int64, d decimal.Decimal) error
	DeleteCacheBalance(ctx context.Context, userID ...int64) error
}

package cache

import (
	"context"

	"github.com/shopspring/decimal"
)

type CacheStorage interface {
	GetBalance(ctx context.Context, userID int64) (decimal *decimal.Decimal, err error)
	UpdateBalance(ctx context.Context, userID int64, d decimal.Decimal) error
	DeleteBalance(ctx context.Context, userID ...int64) error
}

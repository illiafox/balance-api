package balance

import (
	"context"

	"balance-service/app/internal/domain/entity"
	"github.com/shopspring/decimal"
)

type BalanceStorage interface {
	// Balance

	GetBalance(ctx context.Context, userID int64) (decimal decimal.Decimal, err error)
	ChangeBalance(ctx context.Context, userID int64, amount int64, desc string) error

	Transfer(ctx context.Context, fromUserID, toUserID, amount int64, desc string) error

	BlockBalance(ctx context.Context, userID int64, reason string) (err error)
	UnblockBalance(ctx context.Context, userID int64) (err error)

	// Transaction

	GetTransactions(
		ctx context.Context,
		userID, limit, offset int64,
		sort entity.Sort,
	) ([]entity.Transaction, error)
}

type CurrencyStorage interface {
	Get(ctx context.Context, abbreviation string) (decimal.Decimal, error)
}

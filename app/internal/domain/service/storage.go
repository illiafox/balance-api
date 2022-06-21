package service

import (
	"context"

	"balance-service/app/internal/domain/entity"
	"github.com/shopspring/decimal"
)

type BalanceStorage interface {
	// Balance

	GetBalance(ctx context.Context, userID int64) (balance int64, err error)
	ChangeBalance(ctx context.Context, userID, amount int64, desc string) error

	Transfer(ctx context.Context, fromUserID, towUserID, amount int64, desc string) error

	// Transaction

	GetTransactions(
		ctx context.Context,
		userID, limit, offset int64,
		sort entity.Sort,
	) ([]entity.Transaction, error)

	// Owner

	ChangeOwner(ctx context.Context, oldUserID, newUserID int64) error
}

type CurrencyStorage interface {
	Get(ctx context.Context, abbreviation string) (decimal.Decimal, error)
}
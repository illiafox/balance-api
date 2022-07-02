package balance

import (
	"context"

	"balance-service/app/internal/domain/entity"
)

type Service interface {
	Get(ctx context.Context, userID int64, abbr string) (string, error)
	Change(ctx context.Context, userID int64, amount int64, desc string) error
	Transfer(ctx context.Context, fromID int64, toID int64, amount int64, desc string) error
	//
	GetTransactions(ctx context.Context, userID, limit, offset int64, sort string) ([]entity.Transaction, error)
	//
	BlockBalance(ctx context.Context, userID int64, reason string) (err error)
	UnblockBalance(ctx context.Context, userID int64) (err error)
}

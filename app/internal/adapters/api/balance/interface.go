package balance

import (
	"context"

	"balance-service/app/internal/domain/entity"
)

type Service interface {
	Get(ctx context.Context, userID uint64, abbr string) (string, error)
	Change(ctx context.Context, userID uint64, amount int64, desc string) error
	Transfer(ctx context.Context, fromID, toID, amount uint64, desc string) error

	GetTransactions(ctx context.Context, userID, limit, offset uint64, sort string) ([]entity.Transaction, error)

	BlockBalance(ctx context.Context, userID uint64, reason string) (err error)
	UnblockBalance(ctx context.Context, userID uint64) (err error)
}

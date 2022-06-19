package balance

import (
	"context"

	"balance-service/app/internal/domain/entity"
)

type Service interface {
	Get(ctx context.Context, userID int64, abbr string) (string, error)
	Change(ctx context.Context, userID int64, amount int64, desc string) error
	Transfer(ctx context.Context, oldUserID int64, newUserID int64, amount int64, desc string) error
	ChangeOwner(ctx context.Context, oldUserID int64, newUserID int64) error

	GetTransactions(ctx context.Context, userID, limit, offset int64, sort string) ([]entity.Transaction, error)
}

package balance

import (
	"context"
	"fmt"

	"balance-service/app/pkg/errors"
	"github.com/jackc/pgx/v4"
)

func (s balanceStorage) GetBalance(ctx context.Context, userID uint64) (balance int64, err error) {
	// pool.QueryRow() acquires and releases connection automatically
	err = s.pool.QueryRow(ctx, "SELECT balance FROM balance WHERE user_id = $1",
		userID,
	).Scan(&balance)

	if err != nil {
		if err == pgx.ErrNoRows { // no rows -> balance not found
			return -1, fmt.Errorf("balance with user id %d not found", userID)
		}

		return -1, errors.NewInternal(err, "query: get balance")
	}

	return
}

func (balanceStorage) getBalanceForUpdate(ctx context.Context, tx pgx.Tx, userID uint64) (balance uint64, err error) {

	err = tx.QueryRow(ctx, "SELECT balance FROM balance WHERE user_id = $1 FOR UPDATE",
		userID,
	).Scan(&balance)

	if err != nil {
		//nolint:errorlint
		if err == pgx.ErrNoRows { // no rows -> balance not found
			return 0, fmt.Errorf("balance with user id %d not found", userID)
		}

		return 0, errors.NewInternal(err, "query: get balance for update")
	}

	return
}

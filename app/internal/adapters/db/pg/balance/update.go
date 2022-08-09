package balance

import (
	"context"
	"errors"
	"fmt"

	apperrors "balance-service/app/pkg/errors"
	"github.com/jackc/pgx/v4"
)

func (balanceStorage) getBalanceForUpdate(ctx context.Context, tx pgx.Tx, userID int64) (balance int64, err error) {

	err = tx.QueryRow(ctx, "SELECT balance FROM balance WHERE user_id = $1 FOR UPDATE",
		userID,
	).Scan(&balance)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) { // no rows -> balance not found
			return 0, fmt.Errorf("balance with user id %d not found", userID)
		}

		return 0, apperrors.NewInternal(err, "query: get balance for update")
	}

	return
}

func (balanceStorage) updateBalance(ctx context.Context, tx pgx.Tx, userID int64, balance int64) (err error) {
	_, err = tx.Exec(ctx, "UPDATE balance SET balance = $1 WHERE user_id = $2",
		balance, userID,
	)

	if err != nil {
		return apperrors.NewInternal(err, "exec: update balance")
	}

	return nil
}

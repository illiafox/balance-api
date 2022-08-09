package balance

import (
	"context"
	"fmt"

	apperrors "balance-service/app/pkg/errors"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
)

func (s balanceStorage) GetBalance(ctx context.Context, userID int64) (d decimal.Decimal, err error) {
	// pool.QueryRow() acquires and releases connection automatically
	var balance int64

	err = s.pool.QueryRow(ctx, "SELECT balance FROM balance WHERE user_id = $1",
		userID,
	).Scan(&balance)

	if err != nil {
		if err == pgx.ErrNoRows { // no rows -> balance not found
			return d, fmt.Errorf("balance with user id %d not found", userID)
		}

		return d, apperrors.NewInternal(err, "query: get balance")
	}

	d = decimal.New(balance, 0).Shift(-2) // shift 2 decimal places (100 -> 0.01)

	return
}

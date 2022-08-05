package balance

import (
	"context"
	"fmt"

	app_errors "balance-service/app/pkg/errors"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
)

func (s balanceStorage) ChangeBalance(ctx context.Context, userID int64, amount int64, description string) (err error) {
	// acquire connection
	c, err := s.pool.Acquire(ctx)
	if err != nil {
		return app_errors.NewInternal(err, "acquire connection")
	}
	defer c.Release()

	// begin transaction
	tx, err := c.Begin(ctx)
	if err != nil {
		return app_errors.NewInternal(err, "begin transaction")
	}

	defer func() { // defer rollback if error occurs
		if r := tx.Rollback(ctx); r != pgx.ErrTxClosed && r != nil && err == nil {
			err = fmt.Errorf("rollback: %w", r)
		}
	}()

	var balance int64

	// get balance for update
	err = tx.QueryRow(ctx, "SELECT balance FROM balance WHERE user_id = $1 FOR UPDATE",
		userID,
	).Scan(&balance)

	if err != nil {
		//nolint:errorlint
		if err == pgx.ErrNoRows { // no rows -> balance not found -> create new balance
			if amount < 0 { // we can't create new balance with negative amount
				return fmt.Errorf("balance with user id %d not found", userID)
			}
			// check whether user is not blocked
			if err = s.userBlocked(ctx, tx, userID); err != nil {
				return err
			}

			// create new balance
			_, err = tx.Exec(ctx, "INSERT INTO balance (user_id,balance) VALUES ($1,$2)", userID, amount)
			if err != nil {
				return app_errors.NewInternal(err, "exec: create new balance")
			}

		} else { // internal error
			return app_errors.NewInternal(err, "query: get balance for update")
		}
	} else { // if balance found
		balance += amount
		if balance < 0 { // check whether there is enough money to proceed change
			return fmt.Errorf("insufficient funds: missing %s", decimal.NewFromInt(-balance).Shift(-2))
		}

		// update existing balance
		if err = s.updateBalance(ctx, tx, userID, int64(balance)); err != nil {
			return err
		}
	}

	// create record
	_, err = tx.Exec(ctx, `INSERT INTO transaction (to_id,action,description) VALUES ($1,$2,$3)`,
		userID, amount, description,
	)
	if err != nil {
		return app_errors.NewInternal(err, "exec: create record")
	}

	// commit transaction
	err = tx.Commit(ctx)
	if err != nil {
		return app_errors.NewInternal(err, "commit transaction")
	}

	return
}

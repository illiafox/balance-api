package pg

import (
	"context"
	"fmt"

	"balance-service/app/pkg/errors"
	"github.com/jackc/pgx/v4"
)

func (s balanceStorage) BlockBalance(ctx context.Context, userID int64, reason string) (err error) {
	// acquire connection
	c, err := s.pool.Acquire(ctx)
	if err != nil {
		return errors.NewInternal(err, "acquire connection")
	}
	defer c.Release()

	// begin transaction
	tx, err := c.Begin(ctx)
	if err != nil {
		return errors.NewInternal(err, "begin transaction")
	}

	defer func() { // defer rollback if error occurs
		if r := tx.Rollback(ctx); r != pgx.ErrTxClosed && r != nil && err == nil {
			err = fmt.Errorf("rollback: %w", r)
		}
	}()

	var balance int64

	err = tx.QueryRow(ctx, "SELECT balance FROM balance WHERE user_id = $1 FOR UPDATE", userID).Scan(&balance)
	if err != nil {
		//nolint:errorlint
		if err == pgx.ErrNoRows { // no rows -> balance not found
			return fmt.Errorf("balance with user id %d not found", userID)
		}
		return errors.NewInternal(err, "query: get balance for update")
	}

	// delete balance
	_, err = tx.Exec(ctx, "DELETE FROM balance WHERE user_id = $1", userID)
	if err != nil {
		return errors.NewInternal(err, "exec: delete balance")
	}

	// insert blocked balance
	_, err = tx.Exec(ctx, "INSERT INTO block (user_id, balance, reason) VALUES ($1,$2,$3)", userID, balance, reason)
	if err != nil {
		return errors.NewInternal(err, "exec: block balance")
	}

	// commit transaction
	err = tx.Commit(ctx)
	if err != nil {
		return errors.NewInternal(err, "commit transaction")
	}

	return
}

func (s balanceStorage) UnblockBalance(ctx context.Context, userID int64) (err error) {
	// acquire connection
	c, err := s.pool.Acquire(ctx)
	if err != nil {
		return errors.NewInternal(err, "acquire connection")
	}
	defer c.Release()

	// begin transaction
	tx, err := c.Begin(ctx)
	if err != nil {
		return errors.NewInternal(err, "begin transaction")
	}

	defer func() { // defer rollback if error occurs
		if r := tx.Rollback(ctx); r != pgx.ErrTxClosed && r != nil && err == nil {
			err = fmt.Errorf("rollback: %w", r)
		}
	}()

	var balance int64

	err = tx.QueryRow(ctx, "SELECT balance FROM block WHERE user_id = $1 FOR UPDATE", userID).Scan(&balance)
	if err != nil {
		//nolint:errorlint
		if err == pgx.ErrNoRows { // no rows -> balance not found
			return fmt.Errorf("blocked balance with user id %d not found", userID)
		}
		return errors.NewInternal(err, "query: get blocked balance for update")
	}

	// delete blocker balance
	_, err = tx.Exec(ctx, "DELETE FROM block WHERE user_id = $1", userID)
	if err != nil {
		return errors.NewInternal(err, "exec: delete balance")
	}

	// insert balance
	_, err = tx.Exec(ctx, "INSERT INTO balance (user_id, balance) VALUES ($1,$2)", userID, balance)
	if err != nil {
		return errors.NewInternal(err, "exec: block balance")
	}

	// commit transaction
	err = tx.Commit(ctx)
	if err != nil {
		return errors.NewInternal(err, "commit transaction")
	}

	return
}

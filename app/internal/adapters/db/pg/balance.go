package pg

import (
	"context"
	"fmt"

	"balance-service/app/pkg/errors"
	"github.com/jackc/pgx/v4"
)

func (s balanceStorage) GetBalance(ctx context.Context, userID int64) (balance int64, err error) {
	// pool.QueryRow() acquires and releases connection automatically
	err = s.pool.QueryRow(ctx, "SELECT balance FROM balances WHERE user_id = $1", userID).Scan(&balance)
	if err != nil {
		if err == pgx.ErrNoRows { // no rows -> balance not found
			return -1, fmt.Errorf("balance with user id %d not found", userID)
		}

		return -1, errors.NewInternal(err, "query: get balance")
	}

	return
}

func (balanceStorage) getBalanceForUpdate(ctx context.Context, tx pgx.Tx, userID int64) (balance int64, err error) {

	// pool.QueryRow() acquires and releases connection automatically
	err = tx.QueryRow(ctx, "SELECT balance FROM balances WHERE user_id = $1 FOR UPDATE", userID).
		Scan(&balance)
	//
	if err != nil {
		//nolint:errorlint
		if err == pgx.ErrNoRows { // no rows -> balance not found
			return -1, fmt.Errorf("balance with user id %d not found", userID)
		}

		return -1, errors.NewInternal(err, "query: get balance for update")
	}

	return
}

func (s balanceStorage) ChangeBalance(ctx context.Context, userID int64, amount int64, desc string) error {
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
		if r := tx.Rollback(ctx); r != nil && err == nil {
			err = fmt.Errorf("rollback: %w", r)
		}
	}()

	var balance int64

	// get balance
	err = tx.QueryRow(ctx, "SELECT balance FROM balances WHERE user_id = $1 FOR UPDATE", userID).Scan(&balance)

	if err != nil {
		//nolint:errorlint
		if err == pgx.ErrNoRows { // no rows -> balance not found
			if amount < 0 { // we can't create new balance with negative amount
				return fmt.Errorf("balance with user id %d not found", userID)
			}

			// create new balance
			_, err = tx.Exec(ctx, "INSERT INTO balances (user_id,balance) VALUES ($1,$2)", userID, amount)
			if err != nil {
				return errors.NewInternal(err, "exec: create new balance")
			}

		} else { // internal error
			return errors.NewInternal(err, "query: get balance for update")
		}
	} else { // if balance found
		balance += amount
		if balance < 0 { // check whether there is enough money to proceed change
			return fmt.Errorf("insufficient funds: missing %.2f", float64(-balance)/100)
		}

		// update existing balance
		_, err = tx.Exec(ctx, "UPDATE balances SET balance = $1 WHERE user_id = $2", balance, userID)
		if err != nil {
			return errors.NewInternal(err, "exec: update balance")
		}
	}

	// create record
	_, err = tx.Exec(ctx, `INSERT INTO transactions (to_id,action,description) VALUES ($1,$2,$3)`, userID, amount, desc)
	if err != nil {
		return errors.NewInternal(err, "exec: create record")
	}

	// commit transaction
	err = tx.Commit(ctx)
	if err != nil {
		return errors.NewInternal(err, "commit transaction")
	}

	return nil
}

func (s balanceStorage) Transfer(ctx context.Context, fromUserID, toUserID, amount int64, desc string) error {

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
		if r := tx.Rollback(ctx); r != nil && err == nil {
			err = fmt.Errorf("rollback: %w", r)
		}
	}()

	// get sender balance
	sender, err := s.getBalanceForUpdate(ctx, tx, fromUserID)
	if err != nil {
		return err
	}

	// withdraw money from sender
	sender -= amount
	if sender < 0 { // check whether there is enough money to proceed transfer
		return fmt.Errorf("insufficient funds: missing %.2f", float64(-sender)/100)
	}

	//
	var receiver int64
	//

	// get receiver balance
	err = tx.QueryRow(ctx, "SELECT balance FROM balances WHERE user_id = $1 FOR UPDATE", toUserID).
		Scan(&receiver)

	if err != nil {
		//nolint:errorlint
		if err == pgx.ErrNoRows { // no rows -> balance not found
			// create new balance
			_, err = tx.Exec(ctx,
				"INSERT INTO balances (user_id,balance) VALUES ($1,$2)", toUserID, amount,
			)

			if err != nil {
				return errors.NewInternal(err, "exec: create new balance")
			}
		} else { // internal error
			return errors.NewInternal(err, "query: get receiver balance for update")
		}
	} else {
		// update receiver balance
		receiver += amount
		_, err = tx.Exec(ctx, "UPDATE balances SET balance = $1 WHERE user_id = $2", receiver, toUserID)
		if err != nil {
			return errors.NewInternal(err, "exec: update receiver balance")
		}
	}

	// update sender balance
	_, err = tx.Exec(ctx, "UPDATE balances SET balance = $1 WHERE user_id = $2", sender, fromUserID)
	if err != nil {
		return errors.NewInternal(err, "exec: update sender balance")
	}

	// create record
	_, err = tx.Exec(ctx, `INSERT INTO transactions (from_id,to_id,action,description) VALUES ($1,$2,$3,$4)`,
		fromUserID, toUserID, amount, desc)
	if err != nil {
		return errors.NewInternal(err, "exec: create record")
	}

	// commit transaction
	err = tx.Commit(ctx)
	if err != nil {
		return errors.NewInternal(err, "commit transaction")
	}

	return nil
}

package balance

import (
	"context"
	"fmt"

	apperrors "balance-service/app/pkg/errors"
	"github.com/jackc/pgx/v4"
)

func (s balanceStorage) Transfer(ctx context.Context, fromUserID, toUserID, amount int64, desc string) (err error) {

	// acquire connection
	c, err := s.pool.Acquire(ctx)
	if err != nil {
		return apperrors.NewInternal(err, "acquire connection")
	}
	defer c.Release()

	// begin transaction
	tx, err := c.Begin(ctx)
	if err != nil {
		return apperrors.NewInternal(err, "begin transaction")
	}

	defer func() { // defer rollback if error occurs
		if r := tx.Rollback(ctx); r != pgx.ErrTxClosed && r != nil && err == nil {
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
	err = tx.QueryRow(ctx, "SELECT balance FROM balance WHERE user_id = $1 FOR UPDATE",
		toUserID,
	).Scan(&receiver)

	if err != nil {
		//nolint:errorlint
		if err == pgx.ErrNoRows { // no rows -> balance not found
			// check whether user is not blocked
			if err = s.userBlocked(ctx, tx, toUserID); err != nil {
				return err
			}

			// create new balance
			err = s.insertBalanceWithConflict(ctx, tx, toUserID, amount)
			if err != nil {
				return err
			}

		} else { // internal error
			return apperrors.NewInternal(err, "query: get receiver balance for update")
		}
	} else {
		// update receiver balance
		receiver += amount
		if err = s.updateBalance(ctx, tx, toUserID, receiver); err != nil {
			return err
		}
	}

	// update sender balance
	if err = s.updateBalance(ctx, tx, fromUserID, sender); err != nil {
		return err
	}

	record := func(ctx context.Context, tx pgx.Tx, fromUserID, toUserID int64, amount int64, description string) error {
		_, err := tx.Exec(ctx, `INSERT INTO transaction (from_id,to_id,action,description) VALUES ($1,$2,$3,$4)`,
			fromUserID, toUserID, amount, description,
		)
		if err != nil {
			return apperrors.NewInternal(err, "exec: create record")
		}

		return nil
	}

	// // create records

	// deposit money to receiver
	err = record(ctx, tx, fromUserID, toUserID, amount, desc)
	if err != nil {
		return err
	}

	// withdraw money from sender
	err = record(ctx, tx, toUserID, fromUserID, -amount, desc)
	if err != nil {
		return err
	}

	// // commit transaction
	err = tx.Commit(ctx)
	if err != nil {
		return apperrors.NewInternal(err, "commit transaction")
	}

	return
}

package pg

import (
	"context"
	types "database/sql"
	"fmt"
	"strconv"

	"balance-service/app/internal/domain/entity"
	"balance-service/app/pkg/errors"
	"github.com/jackc/pgx/v4"
)

var sorts = map[entity.Sort]string{
	entity.DateASC:  "created_at ASC",
	entity.DateDESC: "created_at DESC",
	entity.SumASC:   "action ASC",
	entity.SumDESC:  "action DESC",
}

func (s *balanceStorage) GetTransactions(
	ctx context.Context,
	userID, limit, offset int64,
	sort entity.Sort,
) ([]entity.Transaction, error) {
	// get ORDER BY clause
	order, ok := sorts[sort]
	if !ok {
		return nil, fmt.Errorf("sort type '%d' not supported", sort)
	}

	// acquire connection
	c, err := s.pool.Acquire(ctx)
	if err != nil {
		return nil, errors.NewInternal(err, "acquire connection")
	}
	defer c.Release()

	var balanceID int64
	// get balance id
	err = c.QueryRow(ctx, "SELECT balance_id FROM balances WHERE user_id = $1", userID).Scan(&balanceID)
	if err != nil {
		//nolint:errorlint
		if err == pgx.ErrNoRows { // no rows -> balance not found
			return nil, fmt.Errorf("balance with user id %d not found", userID)
		}
		return nil, errors.NewInternal(err, "query: get balance")
	}

	// get transactions
	rows, err := c.Query(ctx, "SELECT * FROM transactions WHERE to_id = $1 OR from_id = $1 ORDER BY "+order+" LIMIT $2 OFFSET $3",
		balanceID, limit, offset)
	//
	if err != nil {
		//nolint:errorlint
		if err == pgx.ErrNoRows { // no rows -> no transactions
			return []entity.Transaction{}, nil
		}

		return nil, errors.NewInternal(err, "query: get transactions")
	}
	defer rows.Close()

	var (
		trs  = make([]entity.Transaction, 0, 1)
		t    entity.Transaction
		from types.NullInt64
	)

	for rows.Next() {
		t.FromID = nil // set 'null'

		if err = rows.Scan(&t.ID, &t.ToID, &from, &t.Action, &t.Date, &t.Description); err != nil {
			return nil, errors.NewInternal(err, "scan row")
		}

		// if not null
		if from.Valid {
			t.FromID = []byte(strconv.FormatInt(from.Int64, 10))
		} // else json.RawMessage with 'null'

		trs = append(trs, t)
	}

	return trs, nil
}

package balance

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"balance-service/app/internal/domain/entity"
	apperrors "balance-service/app/pkg/errors"
	sq "github.com/Masterminds/squirrel"
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
		return nil, apperrors.NewInternal(err, "acquire connection")
	}
	defer c.Release()

	// // generate query
	query, args, err := sq.
		Select(
			"transaction_id",
			"to_id",
			"from_id",
			"action",
			"created_at",
			"description",
		).
		From("transaction").
		Where("to_id = $1", userID).
		OrderBy(order).Limit(uint64(limit)).Offset(uint64(offset)).
		ToSql()

	// //

	if err != nil {
		return nil, apperrors.NewInternal(err, "generate sql query")
	}

	// get transactions
	rows, err := c.Query(ctx, query, args...)

	if err != nil {
		//nolint:errorlint
		if err == pgx.ErrNoRows { // no rows -> no transactions
			return nil, nil
		}

		return nil, apperrors.NewInternal(err, "query: get transactions")
	}
	defer rows.Close()

	var (
		trs = make([]entity.Transaction, 0, 1)
		tr  entity.Transaction
		// custom types
		from sql.NullInt64
		t    time.Time
	)

	for rows.Next() {
		if err = rows.Scan(&tr.ID, &tr.ToID, &from, &tr.Action, &t, &tr.Description); err != nil {
			return nil, apperrors.NewInternal(err, "scan row")
		}
		// // custom types

		// from id
		if from.Valid {
			i := from.Int64
			tr.FromID = &i
		} else {
			tr.FromID = nil
		}

		// time
		tr.Date = entity.Time{Time: t}

		// //

		trs = append(trs, tr)
	}

	if err = rows.Err(); err != nil {
		return nil, apperrors.NewInternal(err, "rows")
	}

	return trs, nil
}

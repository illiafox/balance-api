package balance

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"balance-service/app/internal/domain/entity"
	"balance-service/app/pkg/errors"
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
	userID, limit, offset uint64,
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

	// generate query
	query, args, err := sq.Select("*").From("transaction").
		Where("to_id = $1", userID).
		OrderBy(order).Limit(limit).Offset(offset).
		ToSql()

	if err != nil {
		return nil, errors.NewInternal(err, "generate sql query")
	}

	// get transactions
	rows, err := c.Query(ctx, query, args...)

	if err != nil {
		//nolint:errorlint
		if err == pgx.ErrNoRows { // no rows -> no transactions
			return nil, nil
		}

		return nil, errors.NewInternal(err, "query: get transactions")
	}
	defer rows.Close()

	var (
		trs = make([]entity.Transaction, 0, 1)
		tr  entity.Transaction
		// custom types
		t    time.Time
		from sql.NullInt64
	)

	for rows.Next() {
		if err = rows.Scan(&tr.ID, &tr.ToID, &from, &tr.Action, &t, &tr.Description); err != nil {
			return nil, errors.NewInternal(err, "scan row")
		}
		//
		tr.FromID = from.Int64
		tr.Date = entity.Time(t)
		//
		trs = append(trs, tr)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.NewInternal(err, "rows")
	}

	return trs, nil
}

package service

import (
	"context"
	"fmt"

	"balance-service/app/internal/domain/entity"
	"balance-service/app/pkg/errors"
)

func (s *balanceService) GetTransactions(
	ctx context.Context,
	userID, limit, offset int64,
	sort string,
) ([]entity.Transaction, error) {

	var st = entity.DateDESC // default sort

	// find sort type
	if sort != "" {
		var ok bool
		st, ok = entity.Sorts[sort]
		if !ok {
			return nil, fmt.Errorf("sort type '%s' not found", sort)
		}
	}

	trs, err := s.balance.GetTransactions(ctx, userID, limit, offset, st)
	if err != nil {
		if internal, ok := errors.ToInternal(err); ok {
			return nil, internal.Wrap("get transactions")
		}

		return nil, errors.Wrap(err, "get transactions")
	}

	return trs, nil
}

package cache

import (
	"context"
	"fmt"

	app_errors "balance-service/app/pkg/errors"
	"github.com/shopspring/decimal"
)

func (c cacheService) GetBalance(ctx context.Context, userID int64) (d decimal.Decimal, err error) {
	dec, err := c.cache.GetBalance(ctx, userID)
	if err != nil {
		if internal, ok := app_errors.ToInternal(err); ok {
			return d, internal.Wrap("cache")
		}

		return d, fmt.Errorf("cache: %w", err)
	}

	if dec != nil {
		return *dec, nil
	}

	// TODO: remove
	fmt.Println("DB CALL")

	d, err = c.balance.GetBalance(ctx, userID)

	if err == nil {
		err = c.cache.UpdateBalance(ctx, userID, d)
		if err != nil {
			if internal, ok := app_errors.ToInternal(err); ok {
				return d, internal.Wrap("cache")
			}

			return d, fmt.Errorf("update balance: %w", err)
		}
	}

	return

}

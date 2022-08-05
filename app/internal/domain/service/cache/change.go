package cache

import (
	"context"
	"fmt"

	app_errors "balance-service/app/pkg/errors"
)

func (c cacheService) ChangeBalance(ctx context.Context, userID int64, amount int64, desc string) error {
	err := c.cache.DeleteBalance(ctx, userID)
	if err != nil {
		if internal, ok := app_errors.ToInternal(err); ok {
			return internal.Wrap("cache")
		}

		return fmt.Errorf("cache: %w", err)
	}

	return c.balance.ChangeBalance(ctx, userID, amount, desc)
}

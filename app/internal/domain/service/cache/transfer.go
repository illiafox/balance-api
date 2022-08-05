package cache

import (
	"context"
	"fmt"

	app_errors "balance-service/app/pkg/errors"
)

func (c cacheService) Transfer(ctx context.Context, fromUserID, toUserID, amount int64, desc string) error {
	err := c.cache.DeleteBalance(ctx, fromUserID, toUserID)
	if err != nil {
		if internal, ok := app_errors.ToInternal(err); ok {
			return internal.Wrap("cache")
		}

		return fmt.Errorf("cache: %w", err)
	}

	return c.balance.Transfer(ctx, toUserID, fromUserID, amount, desc)
}

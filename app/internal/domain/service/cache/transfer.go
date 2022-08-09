package cache

import (
	"context"
	"fmt"

	apperrors "balance-service/app/pkg/errors"
)

func (c cacheService) Transfer(ctx context.Context, fromUserID, toUserID, amount int64, desc string) error {
	err := c.cache.DeleteCacheBalance(ctx, fromUserID, toUserID)
	if err != nil {
		if internal, ok := apperrors.ToInternal(err); ok {
			return internal.Wrap("cache")
		}

		return fmt.Errorf("cache: %w", err)
	}

	return c.balance.Transfer(ctx, toUserID, fromUserID, amount, desc)
}

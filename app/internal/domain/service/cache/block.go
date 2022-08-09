package cache

import (
	"context"
	"fmt"

	apperrors "balance-service/app/pkg/errors"
)

func (c cacheService) BlockBalance(ctx context.Context, userID int64, reason string) error {
	err := c.cache.DeleteCacheBalance(ctx, userID)
	if err != nil {
		if internal, ok := apperrors.ToInternal(err); ok {
			return internal.Wrap("cache")
		}

		return fmt.Errorf("cache: %w", err)
	}

	return c.balance.BlockBalance(ctx, userID, reason)
}

func (c cacheService) UnblockBalance(ctx context.Context, userID int64) error {
	err := c.cache.DeleteCacheBalance(ctx, userID)
	if err != nil {
		if internal, ok := apperrors.ToInternal(err); ok {
			return internal.Wrap("cache")
		}

		return fmt.Errorf("cache: %w", err)
	}

	return c.balance.UnblockBalance(ctx, userID)
}

package cache

import (
	"context"
	"fmt"

	apperrors "balance-service/app/pkg/errors"
	"github.com/shopspring/decimal"
)

func (c cacheService) GetBalance(ctx context.Context, userID int64) (d decimal.Decimal, err error) {
	dec, err := c.cache.GetBalance(ctx, userID)
	if err != nil {
		if internal, ok := apperrors.ToInternal(err); ok {
			return decimal.Decimal{}, internal.Wrap("cache")
		}

		if err == ErrBalanceNotFound {
			dec, err = c.balance.GetBalance(ctx, userID)

			if err == nil {
				err = c.cache.UpdateBalance(ctx, userID, dec)
				if err != nil {
					if internal, ok := apperrors.ToInternal(err); ok {
						return decimal.Decimal{}, internal.Wrap("cache")
					}

					return decimal.Decimal{}, fmt.Errorf("update balance: %w", err)
				}
			}
		}

		return decimal.Decimal{}, fmt.Errorf("cache: %w", err)
	}

	return dec, nil
}

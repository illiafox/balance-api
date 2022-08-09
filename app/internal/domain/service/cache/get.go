package cache

import (
	"context"
	"fmt"

	apperrors "balance-service/app/pkg/errors"
	"github.com/shopspring/decimal"
)

func (c cacheService) GetBalance(ctx context.Context, userID int64) (d decimal.Decimal, err error) {
	dec, err := c.cache.GetCachedBalance(ctx, userID)
	if err != nil {
		//TODO: remove
		if err == ErrBalanceNotFound {

			dec, err = c.balance.GetBalance(ctx, userID)
			if err != nil {
				return dec, err
			}

			err = c.cache.UpdateCachedBalance(ctx, userID, dec)
			if err != nil {
				if internal, ok := apperrors.ToInternal(err); ok {
					return decimal.Decimal{}, internal.Wrap("cache")
				}

				return decimal.Decimal{}, fmt.Errorf("update balance: %w", err)
			}

			return dec, nil
		}

		if internal, ok := apperrors.ToInternal(err); ok {
			return decimal.Decimal{}, internal.Wrap("cache")
		}

		return decimal.Decimal{}, fmt.Errorf("cache: %w", err)
	}

	return dec, nil
}

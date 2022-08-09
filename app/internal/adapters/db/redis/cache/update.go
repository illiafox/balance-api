package cache

import (
	"context"
	"encoding/json"
	"strconv"

	apperrors "balance-service/app/pkg/errors"
	"github.com/shopspring/decimal"
)

func (c cacheStorage) UpdateCachedBalance(ctx context.Context, userID int64, d decimal.Decimal) error {
	client := c.client.WithContext(ctx)

	id := strconv.FormatInt(userID, 10)

	data, err := json.Marshal(d) // decimal.Decimal has marshal method
	if err != nil {
		return apperrors.NewInternal(err, "marshal decimal")
	}

	err = client.Set(id, data, c.expire).Err()
	if err != nil {
		return apperrors.NewInternal(err, "redis.Set")
	}

	return nil
}

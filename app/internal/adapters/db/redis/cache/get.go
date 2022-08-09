package cache

import (
	"context"
	"encoding/json"
	"strconv"

	"balance-service/app/internal/domain/service/cache"
	apperrors "balance-service/app/pkg/errors"
	"github.com/go-redis/redis"
	"github.com/shopspring/decimal"
)

func (c cacheStorage) GetBalance(ctx context.Context, userID int64) (decimal.Decimal, error) {
	client := c.client.WithContext(ctx)

	id := strconv.FormatInt(userID, 10)

	data, err := client.Get(id).Bytes()
	if err != nil {
		if err == redis.Nil {
			return decimal.Decimal{}, cache.ErrBalanceNotFound
		}
		return decimal.Decimal{}, apperrors.NewInternal(err, "redis.Get")
	}

	var d decimal.Decimal

	err = json.Unmarshal(data, &d) // decimal.Decimal has unmarshal method
	if err != nil {
		return decimal.Decimal{}, apperrors.NewInternal(err, "unmarshal decimal")
	}

	return d, nil
}

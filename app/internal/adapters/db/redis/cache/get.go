package cache

import (
	"context"
	"encoding/json"
	"strconv"

	app_errors "balance-service/app/pkg/errors"
	"github.com/go-redis/redis"
	"github.com/shopspring/decimal"
)

func (c cacheStorage) GetBalance(ctx context.Context, userID int64) (*decimal.Decimal, error) {
	client := c.client.WithContext(ctx)

	id := strconv.FormatInt(userID, 10)

	data, err := client.Get(id).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, app_errors.NewInternal(err, "redis.Get")
	}

	var d decimal.Decimal

	err = json.Unmarshal(data, &d) // decimal.Decimal has unmarshal method
	if err != nil {
		return nil, app_errors.NewInternal(err, "unmarshal decimal")
	}

	return &d, nil
}

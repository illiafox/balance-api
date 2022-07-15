package currency

import (
	"context"
	"fmt"

	"balance-service/app/pkg/errors"
	"github.com/go-redis/redis"
	"github.com/shopspring/decimal"
)

func (s *currencyStorage) Get(ctx context.Context, abbreviation string) (decimal.Decimal, error) {
	c, err := s.client.WithContext(ctx).HGet(s.name, abbreviation).Result()
	if err != nil {
		if err == redis.Nil {
			return decimal.Decimal{}, fmt.Errorf("currency %s not available", abbreviation)
		}

		return decimal.Decimal{}, errors.NewInternal(err, "hget (get data from map)")
	}

	dec, err := decimal.NewFromString(c)
	if err != nil {
		return decimal.Decimal{}, errors.NewInternal(err, "parse decimal")
	}

	return dec, nil
}

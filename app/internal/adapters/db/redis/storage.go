package redis

import (
	"context"
	"fmt"

	"balance-service/app/internal/domain/service"
	"balance-service/app/pkg/errors"
	"github.com/go-redis/redis"
	"github.com/shopspring/decimal"
)

type currencyStorage struct {
	name   string
	client *redis.Client
}

func NewCurrencyStorage(client *redis.Client, name string) service.CurrencyStorage {
	return &currencyStorage{name: name, client: client}
}

func (s *currencyStorage) Get(ctx context.Context, abbreviation string) (decimal.Decimal, error) {
	c, err := s.client.WithContext(ctx).HGet(s.name, abbreviation).Float32()
	if err != nil {
		if err == redis.Nil {
			return decimal.Decimal{}, fmt.Errorf("currency %s not available", abbreviation)
		}

		return decimal.Decimal{}, errors.NewInternal(err, "hget (get data from map)")
	}

	return decimal.NewFromFloat32(c), nil
}

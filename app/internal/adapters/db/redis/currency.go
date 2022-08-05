package redis

import (
	"balance-service/app/internal/adapters/db/redis/currency"
	"balance-service/app/internal/domain/service/balance"
	"github.com/go-redis/redis"
)

func NewCurrencyStorage(client *redis.Client, name string) balance.CurrencyStorage {
	return currency.NewStorage(client, name)
}

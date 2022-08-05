package currency

import (
	"balance-service/app/internal/domain/service/balance"
	"github.com/go-redis/redis"
)

type currencyStorage struct {
	name   string
	client *redis.Client
}

func NewStorage(client *redis.Client, name string) balance.CurrencyStorage {
	return &currencyStorage{name: name, client: client}
}

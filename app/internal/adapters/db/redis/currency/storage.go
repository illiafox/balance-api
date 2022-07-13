package currency

import (
	"balance-service/app/internal/domain/service"
	"github.com/go-redis/redis"
)

type currencyStorage struct {
	name   string
	client *redis.Client
}

func NewStorage(client *redis.Client, name string) service.CurrencyStorage {
	return &currencyStorage{name: name, client: client}
}

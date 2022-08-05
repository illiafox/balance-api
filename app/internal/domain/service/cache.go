package service

import (
	service "balance-service/app/internal/domain/service/balance"
	"balance-service/app/internal/domain/service/cache"
)

func NewCachedBalanceStorage(c cache.CacheStorage, balance service.BalanceStorage) service.BalanceStorage {
	return cache.New(c, balance)
}

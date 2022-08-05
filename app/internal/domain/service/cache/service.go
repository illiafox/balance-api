package cache

import (
	"context"

	"balance-service/app/internal/domain/entity"
	service "balance-service/app/internal/domain/service/balance"
)

type cacheService struct {
	cache   CacheStorage
	balance service.BalanceStorage
}

func (c cacheService) GetTransactions(
	ctx context.Context,
	userID, limit, offset int64,
	sort entity.Sort,
) ([]entity.Transaction, error) {
	return c.balance.GetTransactions(ctx, userID, limit, offset, sort)
}

func New(cache CacheStorage, balance service.BalanceStorage) service.BalanceStorage {
	return cacheService{cache: cache, balance: balance}
}

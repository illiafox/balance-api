package composites

import (
	"balance-service/app/internal/adapters/api/balance"
	"balance-service/app/internal/adapters/db/pg"
	"balance-service/app/internal/adapters/db/redis"
	"balance-service/app/internal/adapters/db/redis/cache"
	"balance-service/app/internal/domain/service"
)

type BalanceComposite struct {
	balance.Service
}

func NewBalanceComposite(pgComposite PgComposite, redisComposite RedisComposite) (*BalanceComposite, error) {
	balanceStorage := pg.NewBalanceStorage(pgComposite.pool)
	cacheStorage := cache.NewCacheStorage(redisComposite.client, redisComposite.cfg.Expire)

	cachedBalanceStorage := service.NewCachedBalanceStorage(cacheStorage, balanceStorage)
	currencyStorage := redis.NewCurrencyStorage(redisComposite.client, redisComposite.hashMap)
	//
	return &BalanceComposite{
		service.NewBalanceService(cachedBalanceStorage, currencyStorage),
	}, nil
}

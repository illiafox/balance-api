package composites

import (
	"balance-service/app/internal/adapters/api"
	pg "balance-service/app/internal/adapters/db/pg"
	"balance-service/app/internal/adapters/db/redis"
	v1 "balance-service/app/internal/controller/http/v1"
	"balance-service/app/internal/domain/service"
)

type BalanceComposite struct {
	Handler api.Handler
}

func NewBalanceComposite(loggerComposite LoggerComposite, pgComposite PgComposite, redisComposite RedisComposite) (*BalanceComposite, error) {
	balanceStorage := pg.NewBalanceStorage(pgComposite.pool)
	currencyStorage := redis.NewCurrencyStorage(redisComposite.client, redisComposite.hashMap)
	//
	balanceService := service.NewBalanceService(balanceStorage, currencyStorage)
	//
	handler := v1.NewHandler(loggerComposite.logger, balanceService)
	//
	return &BalanceComposite{
		Handler: handler,
	}, nil
}

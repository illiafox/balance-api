package composites

import (
	"net/http"

	"balance-service/app/internal/adapters/db/pg"
	"balance-service/app/internal/adapters/db/redis"
	v1 "balance-service/app/internal/controller/http/v1"
	"balance-service/app/internal/domain/service"
	"balance-service/app/pkg/logger"
)

type BalanceComposite struct {
	Handler http.Handler
}

func NewBalanceComposite(logger logger.Logger, pgComposite PgComposite, redisComposite RedisComposite) (*BalanceComposite, error) {
	balanceStorage := pg.NewBalanceStorage(pgComposite.pool)
	currencyStorage := redis.NewCurrencyStorage(redisComposite.client, redisComposite.hashMap)
	//
	balanceService := service.NewBalanceService(balanceStorage, currencyStorage)
	//
	handler := v1.NewHandler(logger, balanceService)
	//
	return &BalanceComposite{
		Handler: handler.Register(),
	}, nil
}

package service

import (
	api "balance-service/app/internal/adapters/api/balance"
	balance "balance-service/app/internal/domain/service/balance"
)

func NewBalanceService(storage balance.BalanceStorage, currency balance.CurrencyStorage) api.Service {
	return balance.New(storage, currency)
}

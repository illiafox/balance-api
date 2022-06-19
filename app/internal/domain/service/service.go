package service

import (
	"balance-service/app/internal/adapters/api/balance"
)

type balanceService struct {
	balance  BalanceStorage
	currency CurrencyStorage
}

func NewBalanceService(balance BalanceStorage, currency CurrencyStorage) balance.Service {
	return &balanceService{balance: balance, currency: currency}
}

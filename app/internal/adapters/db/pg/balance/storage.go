package balance

import (
	"balance-service/app/internal/domain/service/balance"
	"github.com/jackc/pgx/v4/pgxpool"
)

type balanceStorage struct {
	pool *pgxpool.Pool
}

func NewStorage(pool *pgxpool.Pool) balance.BalanceStorage {
	return &balanceStorage{pool: pool}
}

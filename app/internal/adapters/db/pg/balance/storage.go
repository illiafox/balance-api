package balance

import (
	"balance-service/app/internal/domain/service"
	"github.com/jackc/pgx/v4/pgxpool"
)

type balanceStorage struct {
	pool *pgxpool.Pool
}

func NewStorage(pool *pgxpool.Pool) service.BalanceStorage {
	return &balanceStorage{pool: pool}
}

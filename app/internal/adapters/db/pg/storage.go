package pg

import (
	"balance-service/app/internal/domain/service"
	"github.com/jackc/pgx/v4/pgxpool"
)

type balanceStorage struct {
	pool *pgxpool.Pool
}

func NewBalanceStorage(pool *pgxpool.Pool) service.BalanceStorage {
	return &balanceStorage{pool: pool}
}

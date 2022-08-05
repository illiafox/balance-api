package pg

import (
	"balance-service/app/internal/adapters/db/pg/balance"
	service "balance-service/app/internal/domain/service/balance"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewBalanceStorage(pool *pgxpool.Pool) service.BalanceStorage {
	return balance.NewStorage(pool)
}

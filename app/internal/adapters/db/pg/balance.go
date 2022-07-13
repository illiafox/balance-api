package pg

import (
	"balance-service/app/internal/adapters/db/pg/balance"
	"balance-service/app/internal/domain/service"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewBalanceStorage(pool *pgxpool.Pool) service.BalanceStorage {
	return balance.NewStorage(pool)
}

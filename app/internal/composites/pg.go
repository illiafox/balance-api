package composites

import (
	"context"
	"fmt"

	"balance-service/app/internal/config"
	"balance-service/app/pkg/client/pg"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PgComposite struct {
	pool *pgxpool.Pool
}

func (pg PgComposite) Close() error {
	pg.pool.Close()
	return nil
}

func NewPgComposite(ctx context.Context, cfg config.Postgres) (PgComposite, error) {
	var composite PgComposite

	pool, err := pg.NewPool(ctx, cfg.User, cfg.Pass, cfg.Database, cfg.IP, cfg.Port, cfg.Protocol)
	if err != nil {
		return composite, fmt.Errorf("create pool: %w", err)
	}

	composite.pool = pool
	return composite, nil
}

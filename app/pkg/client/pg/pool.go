package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPool(ctx context.Context, user string, pass string, db string, ip string, port int) (*pgxpool.Pool, error) {
	return pgxpool.Connect(
		ctx,
		fmt.Sprintf("postgres://%s:%s@%v:%v/%v?sslmode=disable",
			user,
			pass,
			ip,
			port,
			db,
		),
	)
}

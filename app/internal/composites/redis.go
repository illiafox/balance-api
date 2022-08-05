package composites

import (
	"context"
	"fmt"

	"balance-service/app/internal/config"
	"balance-service/app/pkg/client/redis"
	rdb "github.com/go-redis/redis"
)

type RedisComposite struct {
	hashMap string
	client  *rdb.Client
	cfg     config.Redis
}

func (r RedisComposite) Close() error {
	return r.client.Close()
}

func NewRedisComposite(ctx context.Context, cfg config.Redis) (RedisComposite, error) {
	var composite = RedisComposite{
		hashMap: cfg.HashMap,
		cfg:     cfg,
	}

	client, err := redis.New(ctx, fmt.Sprintf("%s:%d", cfg.IP, cfg.Port), cfg.Pass, cfg.DB)
	if err != nil {
		return composite, fmt.Errorf("create client: %w", err)
	}

	composite.client = client
	return composite, nil
}

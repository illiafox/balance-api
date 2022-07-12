package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

func New(ctx context.Context, address string, pass string, db int) (*redis.Client, error) {
	options := &redis.Options{
		Addr:     address,
		Password: pass,
		//
		DB: db,
	}

	client := redis.NewClient(options)

	if err := client.WithContext(ctx).
		Ping().Err(); err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	return client, nil
}

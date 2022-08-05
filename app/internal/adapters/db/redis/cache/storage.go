package cache

import (
	"time"

	"balance-service/app/internal/domain/service/cache"
	"github.com/go-redis/redis"
)

type cacheStorage struct {
	client *redis.Client
	expire time.Duration
}

func NewCacheStorage(client *redis.Client, expire time.Duration) cache.CacheStorage {
	return cacheStorage{client: client, expire: expire}
}

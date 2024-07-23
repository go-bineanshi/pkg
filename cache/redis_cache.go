package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisCache(client *redis.Client, ttl time.Duration) *RedisCache {
	return &RedisCache{
		client: client,
		ttl:    ttl,
	}
}

func (rc *RedisCache) Set(key string, value interface{}) error {
	return rc.client.Set(context.Background(), key, value, rc.ttl).Err()
}

func (rc *RedisCache) Get(key string) (string, error) {
	return rc.client.Get(context.Background(), key).Result()
}
func (rc *RedisCache) Delete(key string) error {
	return rc.client.Del(context.Background(), key).Err()
}

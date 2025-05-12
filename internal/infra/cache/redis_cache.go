package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	DefaultExpiration = 10 * time.Minute
)

type ICacheRepository interface {
	Set(ctx context.Context, key string, value interface{}, exp time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}

type Cache struct {
	redis *redis.Client
}

func New(redis *redis.Client) ICacheRepository {
	return &Cache{redis: redis}
}

func (r *Cache) Set(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	err := r.redis.Set(ctx, key, value, exp).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Cache) Get(ctx context.Context, key string) (string, error) {
	val, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (r *Cache) Delete(ctx context.Context, key string) error {
	err := r.redis.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}

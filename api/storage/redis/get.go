package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func (rc *RedisClient) Get(ctx context.Context, key string) (string, error) {
	value, err := rc.client.Get(ctx, key).Result()
	if err != redis.Nil && err != nil {
		return "", fmt.Errorf("failed to get value from Redis: %w", err)
	}

	return value, nil
}

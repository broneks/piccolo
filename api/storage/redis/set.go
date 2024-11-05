package redis

import (
	"context"
	"time"
)

func (rc *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := rc.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rc *RedisClient) SetForever(ctx context.Context, key string, value interface{}) error {
	return rc.Set(ctx, key, value, 0)
}

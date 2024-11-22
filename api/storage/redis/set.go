package redis

import (
	"context"
	"time"
)

func (rc *RedisClient) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	err := rc.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

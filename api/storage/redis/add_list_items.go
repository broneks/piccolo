package redis

import (
	"context"
	"fmt"
)

func (rc *RedisClient) AddListItems(ctx context.Context, key string, values ...any) error {
	err := rc.client.SAdd(ctx, key, values)
	if err != nil {
		return fmt.Errorf("failed to add list items to Redis: %v", err)
	}

	return nil
}

package redis

import (
	"context"
	"fmt"
)

func (rc *RedisClient) IsListItem(ctx context.Context, key string, value any) (bool, error) {
	isMember, err := rc.client.SIsMember(ctx, key, value).Result()
	if err != nil {
		return false, fmt.Errorf("failed to check list item in Redis: %w", err)
	}

	return isMember, nil
}

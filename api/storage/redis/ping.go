package redis

import "context"

func (rc *RedisClient) Ping(ctx context.Context) error {
	return rc.client.Ping(ctx).Err()
}

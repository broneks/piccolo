package banhammerservice

import (
	"context"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

func (svc *BanHammerService) IsBanned(ctx context.Context, ip string) (bool, time.Duration) {
	key := banKeyPrefix + ip

	ttl, err := svc.rdb.TTL(ctx, key).Result()
	if err != nil && err != redis.Nil {
		slog.Error("Error checking ban status:", "error", err)
	}

	return ttl > 0, ttl
}

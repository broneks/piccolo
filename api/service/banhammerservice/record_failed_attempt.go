package banhammerservice

import (
	"context"
	"fmt"
	"log/slog"
)

func (svc *BanHammerService) RecordFailedAttempt(ctx context.Context, ip string) error {
	attemptsKey := attemptKeyPrefix + ip
	banKey := banKeyPrefix + ip

	// Increment the failed attempts count atomically
	attempts, err := svc.rdb.Incr(ctx, attemptsKey).Result()
	if err != nil {
		return err
	}

	// Set expiration for the tracking key if it’s the first attempt
	if attempts == 1 {
		svc.rdb.Expire(ctx, attemptsKey, trackingTime)
	}

	// Check if the IP should be banned
	if attempts >= maxAttempts {
		err := svc.rdb.Set(ctx, banKey, "1", banDuration).Err() // Ban the IP
		if err != nil {
			return err
		}
		// Cleanup the attempts key
		svc.rdb.Del(ctx, attemptsKey)
		slog.Info(fmt.Sprintf("IP %s has been banned for %s\n", ip, banDuration))
	}

	return nil
}

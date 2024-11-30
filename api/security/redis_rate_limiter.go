package security

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	redisrate "github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

type RedisRateLimiter struct {
	*redisrate.Limiter
}

const rateRequest = "rate_request_%s"

func NewRedisRateLimiter() *RedisRateLimiter {
	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatalf("cannot create redis connection: %v", err)
	}

	rdb := redis.NewClient(opts)

	return &RedisRateLimiter{
		redisrate.NewLimiter(rdb),
	}
}

func (rl *RedisRateLimiter) Limit(ctx context.Context, ip string) (*redisrate.Result, error) {
	if ip == "" {
		return nil, errors.New("IP is not provided")
	}

	return rl.Allow(ctx, fmt.Sprintf(rateRequest, ip), redisrate.PerSecond(5))
}

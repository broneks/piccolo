package redis

import (
	"os"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewClient() *RedisClient {
	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opts)

	return &RedisClient{rdb}
}

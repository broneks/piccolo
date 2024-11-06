package redis

import (
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
}

func NewClient() *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	return &RedisClient{rdb}
}

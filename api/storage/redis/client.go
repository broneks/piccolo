package redis

import (
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
}

func NewClient() (*RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	client := &RedisClient{rdb}

	return client, nil
}

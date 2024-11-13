package redis

import (
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewClient() *RedisClient {
	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatalf("cannot create redis connection: %v", err)
	}

	rdb := redis.NewClient(opts)

	return &RedisClient{rdb}
}

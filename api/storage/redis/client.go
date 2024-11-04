package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

func Get(c echo.Context, key string) (string, error) {
	reqCtx := c.Request().Context()
	rdb := c.Get("redisClient").(*redis.Client)

	val, err := rdb.Get(reqCtx, key).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get value from Redis: %w", err)
	}

	return val, nil
}

func Set(c echo.Context, key string, value string) error {
	reqCtx := c.Request().Context()
	rdb := c.Get("redisClient").(*redis.Client)

	err := rdb.Set(reqCtx, key, value, 0).Err() // 0 means no expiration
	if err != nil {
		return fmt.Errorf("failed to set value in Redis: %w", err)
	}

	return nil
}

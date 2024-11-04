package model

import (
	"context"
	"fmt"
	"piccolo/api/storage/wasabi"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

type Photo struct {
	Id          string
	Location    string
	Filename    string
	FileSize    int
	ContentType string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (p *Photo) GetUrl(c echo.Context) string {
	val, _ := getFromRedis(c, p.Filename)
	if val == "" {
		fmt.Println("no cache found for url image.")
		url, _ := wasabi.GetPresignedUrl(context.Background(), p.Filename)
		setInRedis(c, p.Filename, url)
		return url
	} else {
		fmt.Println("found cached url for image")
		return val
	}
}

func getFromRedis(c echo.Context, key string) (string, error) {
	reqCtx := c.Request().Context()
	rdb := c.Get("redisClient").(*redis.Client)

	val, err := rdb.Get(reqCtx, key).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get value from Redis: %w", err)
	}

	return val, nil
}

func setInRedis(c echo.Context, key string, value string) error {
	reqCtx := c.Request().Context()
	rdb := c.Get("redisClient").(*redis.Client)

	err := rdb.Set(reqCtx, key, value, 0).Err() // 0 means no expiration
	if err != nil {
		return fmt.Errorf("failed to set value in Redis: %w", err)
	}

	return nil
}

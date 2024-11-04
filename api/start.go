package api

import (
	"context"
	"piccolo/api/upload"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type contextKey string

const redisKey contextKey = "redisClient"

var ctx = context.Background()

func Start() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379", // Address of the Redis container
	})

	e := echo.New()

	// Middleware to set the Redis client in the context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(string(redisKey), rdb) // Store the Redis client in the context
			return next(c)
		}
	})

	e.Use(Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())

	e.Static("/", "public")

	upload.Router(e)

	e.HideBanner = true
	e.Logger.Fatal(e.Start(":8001"))
}

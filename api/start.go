package api

import (
	"context"
	"fmt"
	"log/slog"
	"piccolo/api/middleware"
	"piccolo/api/modules"
	"piccolo/api/shared"
	"piccolo/api/storage/pg"
	"piccolo/api/storage/redis"
	"piccolo/api/storage/wasabi"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Start() {
	var err error

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.RequestID())
	e.Use(echoMiddleware.Secure())
	e.Static("/", "public")

	logger := slog.Default()

	db, err := pg.NewClient(context.Background())
	if err != nil {
		logger.Error(fmt.Sprintf("Cannot load postgres db: %v", err))
		return
	}

	redis, err := redis.NewClient()
	if err != nil {
		logger.Error(fmt.Sprintf("Cannot load redis db: %v", err))
		return
	}

	wasabi, err := wasabi.NewClient(context.Background())
	if err != nil {
		logger.Error(fmt.Sprintf("Cannot load wasabi db: %v", err))
		return
	}

	server := &shared.Server{
		Logger: logger,
		DB:     db,
		Redis:  redis,
		Wasabi: wasabi,
	}

	modules.Routes(e, server)

	e.HideBanner = true
	e.Logger.Fatal(e.Start(":8001"))
}

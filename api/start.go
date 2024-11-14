package api

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"piccolo/api/middleware"
	"piccolo/api/modules"
	"piccolo/api/shared"
	"piccolo/api/storage/backblaze"
	"piccolo/api/storage/pg"
	"piccolo/api/storage/redis"
	"piccolo/api/util"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Start() {
	env := os.Getenv("ENV")

	var err error

	e := echo.New()

	e.Validator = shared.NewValidator()

	// custom
	e.Use(middleware.Logger())
	e.Use(middleware.SetUserData())

	// echo built-in
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.RequestID())
	e.Use(echoMiddleware.Secure())

	e.Static("/", "static")

	logger := slog.Default()

	dbClient, err := pg.NewClient(context.Background())
	if err != nil {
		logger.Error(fmt.Sprintf("Cannot create database client: %v", err.Error()))
		os.Exit(1)
	}

	redisClient := redis.NewClient()

	backblazeClient, err := backblaze.NewClient(context.Background())
	if err != nil {
		logger.Error(fmt.Sprintf("Cannot create backblaze client: %v", err.Error()))
		os.Exit(1)
	}

	server := &shared.Server{
		Logger:        logger,
		DB:            dbClient,
		Cache:         redisClient,
		ObjectStorage: backblazeClient,
	}

	g := e.Group("/api")
	modules.Routes(g, server)

	if env == "local" {
		util.ListAllRoutes(e)
	}

	e.HideBanner = true
	e.Logger.Fatal(e.Start("0.0.0.0:8001"))
}

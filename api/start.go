package api

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"piccolo/api/middleware"
	"piccolo/api/module"
	"piccolo/api/page"
	"piccolo/api/storage/backblaze"
	"piccolo/api/storage/pg"
	"piccolo/api/storage/redis"
	"piccolo/api/types"
	"piccolo/api/util"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func newServer(ctx context.Context) *types.Server {
	var err error

	logger := slog.Default()

	dbClient, err := pg.NewClient(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("Cannot create database client: %v", err.Error()))
		os.Exit(1)
	}

	redisClient := redis.NewClient()

	backblazeClient, err := backblaze.NewClient(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("Cannot create backblaze client: %v", err.Error()))
		os.Exit(1)
	}

	return &types.Server{
		Logger:        logger,
		DB:            dbClient,
		Cache:         redisClient,
		ObjectStorage: backblazeClient,
	}
}

func Start() {
	env := os.Getenv("ENV")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.RequestID())
	e.Use(echoMiddleware.Secure())

	e.Static("/", "static")

	e.IPExtractor = echo.ExtractIPDirect()
	e.Validator = util.NewValidator()
	e.Renderer = util.NewTemplateRenderer("templates/*.html")
	e.HTTPErrorHandler = httpErrorHandler

	server := newServer(context.Background())

	page.Routes(e, server)

	module.Routes(e.Group("/api"), server)

	if env == "local" {
		util.ListAllRoutes(e)
	}

	e.HideBanner = true
	e.Logger.Fatal(e.Start("0.0.0.0:8001"))
}

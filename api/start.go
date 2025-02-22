package api

import (
	"context"
	"log/slog"
	"os"
	"piccolo/api/helper"
	"piccolo/api/mailer"
	"piccolo/api/middleware"
	"piccolo/api/page"
	"piccolo/api/resource"
	"piccolo/api/service/rendererservice"
	"piccolo/api/service/validatorservice"
	"piccolo/api/storage/backblaze"
	"piccolo/api/storage/pg"
	"piccolo/api/storage/redis"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func newServer(ctx context.Context) *types.Server {
	var err error

	dbClient, err := pg.NewClient(ctx)
	if err != nil {
		slog.Error("cannot create database client", "err", err)
		os.Exit(1)
	}

	redisClient := redis.NewClient()

	backblazeClient, err := backblaze.NewClient(ctx)
	if err != nil {
		slog.Error("cannot create backblaze client", "err", err)
		os.Exit(1)
	}

	mailerClient := mailer.New()

	return &types.Server{
		Mailer:        mailerClient,
		DB:            dbClient,
		Cache:         redisClient,
		ObjectStorage: backblazeClient,
	}
}

func Start() {
	env := os.Getenv("ENV")

	e := echo.New()

	e.IPExtractor = echo.ExtractIPDirect()

	e.Pre(middleware.HttpsRedirect())
	e.Pre(middleware.HttpsNonWWWRedirect())
	e.Pre(echoMiddleware.RemoveTrailingSlash())

	e.Use(middleware.RateLimiter())
	e.Use(middleware.Logger())
	e.Use(middleware.CacheControl())
	e.Use(middleware.Cors())
	e.Use(middleware.Csrf())
	e.Use(middleware.Secure())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.RequestID())

	e.Static("/", "static")

	e.Validator = validatorservice.New()
	e.Renderer = rendererservice.New("templates/*.html")
	e.HTTPErrorHandler = httpErrorHandler

	server := newServer(context.Background())

	page.Routes(e, server)
	resource.Routes(e.Group("/api"), server)

	if env == "local" {
		helper.ListAllRoutes(e)
	}

	e.HideBanner = true
	e.Logger.Fatal(e.Start("0.0.0.0:8000"))
}

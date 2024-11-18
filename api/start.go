package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"piccolo/api/middleware"
	"piccolo/api/modules"
	"piccolo/api/storage/backblaze"
	"piccolo/api/storage/pg"
	"piccolo/api/storage/redis"
	"piccolo/api/types"
	"piccolo/api/util"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Start() {
	env := os.Getenv("ENV")

	var err error

	e := echo.New()

	e.IPExtractor = echo.ExtractIPDirect()
	e.Validator = util.NewValidator()

	e.Use(middleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.RequestID())
	e.Use(echoMiddleware.Secure())

	logger := slog.Default()

	e.Static("/", "static")
	e.Renderer = util.NewTemplateRenderer("templates/*.html")
	e.HTTPErrorHandler = httpErrorHandler

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

	server := &types.Server{
		Logger:        logger,
		DB:            dbClient,
		Cache:         redisClient,
		ObjectStorage: backblazeClient,
	}

	// TODO remove
	e.GET("/test", func(c echo.Context) error {
		return c.Render(http.StatusOK, "test.html", map[string]any{
			"name": "Dolly!",
		})
	})

	g := e.Group("/api")
	modules.Routes(g, server)

	if env == "local" {
		util.ListAllRoutes(e)
	}

	e.HideBanner = true
	e.Logger.Fatal(e.Start("0.0.0.0:8001"))
}

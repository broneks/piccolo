package middleware

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/lmittmann/tint"
)

func Logger() echo.MiddlewareFunc {
	w := os.Stderr

	logger := slog.New(tint.NewHandler(w, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.Kitchen,
	}))

	slog.SetDefault(logger)

	return echoMiddleware.RequestLoggerWithConfig(echoMiddleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v echoMiddleware.RequestLoggerValues) error {
			ctx := c.Request().Context()

			method := v.Method
			if method == "" {
				method = "GET"
			}

			if v.Error == nil {
				logger.LogAttrs(
					ctx,
					slog.LevelInfo,
					"request",
					slog.String("uri", fmt.Sprintf("%s %s", method, v.URI)),
					slog.Int("status", v.Status),
				)
			} else {
				logger.LogAttrs(
					ctx,
					slog.LevelError,
					"request_error",
					slog.String("uri", fmt.Sprintf("%s %s", method, v.URI)),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	})
}

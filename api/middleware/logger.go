package middleware

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Logger() echo.MiddlewareFunc {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	return echoMiddleware.RequestLoggerWithConfig(echoMiddleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v echoMiddleware.RequestLoggerValues) error {
			method := v.Method
			if method == "" {
				method = "GET"
			}

			if v.Error == nil {
				logger.LogAttrs(
					c.Request().Context(),
					slog.LevelInfo,
					"request",
					slog.String("uri", fmt.Sprintf("%s %s", method, v.URI)),
					slog.Int("status", v.Status),
				)
			} else {
				logger.LogAttrs(
					c.Request().Context(),
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

package middleware

import (
	"log/slog"
	"net/http"
	"piccolo/api/security"

	"github.com/labstack/echo/v4"
)

func RateLimiter() echo.MiddlewareFunc {
	limiter := security.NewRedisRateLimiter()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			ip := c.RealIP()

			res, err := limiter.Limit(ctx, ip)
			if err != nil {
				slog.Debug(err.Error())
			}

			if res.Allowed <= 0 {
				return echo.NewHTTPError(http.StatusTooManyRequests)
			}

			return next(c)
		}
	}
}

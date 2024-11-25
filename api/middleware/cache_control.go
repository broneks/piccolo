package middleware

import "github.com/labstack/echo/v4"

func CacheControl() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Cache-Control", "private, max-age=3600") // 1 hour in seconds

			return next(c)
		}
	}
}

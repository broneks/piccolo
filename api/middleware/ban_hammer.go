package middleware

import (
	"fmt"
	"net/http"
	"piccolo/api/service/banhammerservice"

	"github.com/labstack/echo/v4"
)

func BanHammer(banHammerService *banhammerservice.BanHammerService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			ip := c.RealIP()

			banned, ttl := banHammerService.IsBanned(ctx, ip)
			if banned {
				return echo.NewHTTPError(http.StatusForbidden, fmt.Sprintf("You are banned for %s", ttl))
			}

			return next(c)
		}
	}
}

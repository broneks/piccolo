package middleware

import (
	"net/http"
	"piccolo/api/consts"
	"piccolo/api/service/jwtservice"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func CanResetPassword(server *types.Server) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()

			var tokenString string

			tokenString = c.QueryParam("token")

			// retry a different way
			if tokenString == "" {
				tokenString = c.FormValue("token")
			}

			if tokenString == "" {
				return echo.NewHTTPError(http.StatusForbidden)
			}

			isBlacklisted, err := server.Cache.IsListItem(ctx, consts.ResetPasswordTokenBlacklistKey, tokenString)
			if err != nil {
				server.Logger.Error(err.Error())
			}
			if isBlacklisted {
				return echo.NewHTTPError(http.StatusForbidden)
			}

			isAuthenticated := jwtservice.VerifyToken(tokenString)
			if !isAuthenticated {
				return echo.NewHTTPError(http.StatusForbidden)
			}

			return next(c)
		}
	}
}

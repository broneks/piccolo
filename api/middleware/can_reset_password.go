package middleware

import (
	"log/slog"
	"net/http"
	"piccolo/api/jwtoken"

	"github.com/labstack/echo/v4"
)

func CanResetPassword() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var tokenString string

			tokenString = c.QueryParam("token")

			// retry a different way
			if tokenString == "" {
				tokenString = c.FormValue("token")
			}

			slog.Debug(tokenString)

			isAuthenticated := jwtoken.VerifyToken(tokenString)
			if !isAuthenticated {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			return next(c)
		}
	}
}

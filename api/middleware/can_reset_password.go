package middleware

import (
	"net/http"
	"piccolo/api/jwtoken"

	"github.com/labstack/echo/v4"
)

func CanResetPassword() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.QueryParam("token")

			isAuthenticated := jwtoken.VerifyToken(tokenString)
			if !isAuthenticated {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			return next(c)
		}
	}
}

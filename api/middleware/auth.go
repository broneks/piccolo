package middleware

import (
	"log/slog"
	"net/http"
	"piccolo/api/jwtoken"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			tokenString, err := jwtoken.ExtractTokenString(authHeader)
			if err != nil {
				slog.Error(err.Error())
			}

			isAuthenticated := jwtoken.VerifyToken(tokenString)
			if !isAuthenticated {
				return c.JSON(http.StatusUnauthorized, types.SuccessRes{
					Success: false,
					Message: "Unauthenticated: Please login to continue.",
				})
			}

			return next(c)
		}
	}
}

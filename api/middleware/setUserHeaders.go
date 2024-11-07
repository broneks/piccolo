package middleware

import (
	"net/http"
	"piccolo/api/jwtoken"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func SetUserHeaders() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString, err := jwtoken.ExtractTokenString(c.Request().Header.Get("Authorization"))
			if err != nil {
				log.Error(err)
			}

			userId := jwtoken.GetUserId(tokenString)
			if userId == "" {
				log.Error("User Id not found in token")
			}

			userEmail := jwtoken.GetUserEmail(tokenString)
			if userEmail == "" {
				log.Error("User email not found in token")
			}

			if err != nil || userId == "" || userEmail == "" {
				return c.JSON(http.StatusUnauthorized, "User authentication failed.")
			}

			c.Response().Header().Set("X-User-Id", userId)
			c.Response().Header().Set("X-User-Email", userEmail)

			return next(c)
		}
	}
}

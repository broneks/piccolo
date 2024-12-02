package middleware

import (
	"log/slog"
	"piccolo/api/service/jwtservice"

	"github.com/labstack/echo/v4"
)

type UserData struct {
	Id    string
	Email string
}

func SetUserData() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString, err := getAccesssTokenString(c)
			if err != nil {
				slog.Error("Error setting user data", "err", err)
			}

			if tokenString != "" {
				userId := jwtservice.GetUserId(tokenString)
				userEmail := jwtservice.GetUserEmail(tokenString)

				c.Set("userId", userId)
				c.Set("userEmail", userEmail)
			}

			return next(c)
		}
	}
}

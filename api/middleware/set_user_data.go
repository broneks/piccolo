package middleware

import (
	"log"
	"piccolo/api/jwtoken"

	"github.com/labstack/echo/v4"
)

type UserData struct {
	Id    string
	Email string
}

func SetUserData() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString, err := jwtoken.ExtractTokenString(c.Request().Header.Get("Authorization"))
			if err != nil {
				log.Println("error setting user data", err.Error())
			}

			if tokenString != "" {
				userId := jwtoken.GetUserId(tokenString)
				userEmail := jwtoken.GetUserEmail(tokenString)

				c.Set("userId", userId)
				c.Set("userEmail", userEmail)
			}

			return next(c)
		}
	}
}

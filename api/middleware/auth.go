package middleware

import (
	"log/slog"
	"net/http"
	"piccolo/api/jwtoken"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func getAccesssTokenString(c echo.Context) (string, error) {
	var err error

	// first try getting the jwt via an http-only cookie
	tokenCookie, err := c.Cookie("piccolo-access-token")
	if err != nil {
		slog.Debug(err.Error())
	}
	if tokenCookie != nil {
		return tokenCookie.Value, nil
	}

	// fallback to using the auth header
	tokenString, err := jwtoken.ExtractTokenString(c.Request().Header.Get("Authorization"))
	if err != nil {
		slog.Debug(err.Error())
		return "", err
	}

	return tokenString, nil
}

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString, err := getAccesssTokenString(c)
			if err != nil {
				slog.Error(err.Error())
			}

			isAuthenticated := jwtoken.VerifyToken(tokenString)
			if !isAuthenticated {
				return c.JSON(http.StatusUnauthorized, types.SuccessRes{
					Success: false,
					Message: "Unauthorized: Please login to continue.",
				})
			}

			return next(c)
		}
	}
}

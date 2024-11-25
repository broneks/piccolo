package middleware

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Csrf() echo.MiddlewareFunc {
	return echoMiddleware.CSRFWithConfig(echoMiddleware.CSRFConfig{
		Skipper:        Skipper,
		TokenLookup:    "form:csrf",
		CookieName:     "_csrf",
		CookieSecure:   true,
		CookieHTTPOnly: true,
	})
}

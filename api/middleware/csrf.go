package middleware

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Csrf() echo.MiddlewareFunc {
	return echoMiddleware.CSRFWithConfig(echoMiddleware.CSRFConfig{
		Skipper:        Skipper,
		TokenLookup:    "cookie:_csrf",
		CookiePath:     "/",
		CookieDomain:   "piccolo.pics",
		CookieSecure:   true,
		CookieHTTPOnly: true,
	})
}

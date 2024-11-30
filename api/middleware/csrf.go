package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func csrfSkipper(c echo.Context) bool {
	if Skipper(c) {
		return true
	}

	url := c.Request().URL.String()

	return strings.HasPrefix(url, "/api")
}

func Csrf() echo.MiddlewareFunc {
	return echoMiddleware.CSRFWithConfig(echoMiddleware.CSRFConfig{
		Skipper:        csrfSkipper,
		TokenLookup:    "form:csrf",
		CookieName:     "_csrf",
		CookieSecure:   true,
		CookieHTTPOnly: true,
	})
}

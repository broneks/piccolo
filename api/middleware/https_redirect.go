package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func HttpsRedirect() echo.MiddlewareFunc {
	return echoMiddleware.HTTPSRedirectWithConfig(echoMiddleware.RedirectConfig{
		Skipper: Skipper,
		Code:    http.StatusMovedPermanently,
	})
}

func HttpsNonWWWRedirect() echo.MiddlewareFunc {
	return echoMiddleware.HTTPSNonWWWRedirectWithConfig(echoMiddleware.RedirectConfig{
		Skipper: Skipper,
		Code:    http.StatusMovedPermanently,
	})
}

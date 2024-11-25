package middleware

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func allowOrigin(origin string) (bool, error) {
	return regexp.MatchString(`^https:\/\/piccolo.pics$`, origin)
}

func Cors() echo.MiddlewareFunc {
	return echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		Skipper:         Skipper,
		AllowOriginFunc: allowOrigin,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPatch,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
		},
	})
}

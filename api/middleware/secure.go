package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func generateNonce() string {
	nonce := make([]byte, 16)

	_, err := rand.Read(nonce)
	if err != nil {
		panic("failed to generate nonce")
	}

	return base64.StdEncoding.EncodeToString(nonce)
}

func Secure() echo.MiddlewareFunc {
	nonce := generateNonce()

	secureMiddlewareFunc := echoMiddleware.SecureWithConfig(echoMiddleware.SecureConfig{
		Skipper:            echoMiddleware.DefaultSkipper,
		XSSProtection:      "1; mode=block",
		ContentTypeNosniff: "nosniff",
		ContentSecurityPolicy: fmt.Sprintf(`
			default-src 'none';
			script-src 'self' 'nonce-%s' https://cdn.tailwindcss.com;
			connect-src 'self';
			img-src 'self' https://f005.backblazeb2.com data:;
			style-src 'self' 'unsafe-inline';
			font-src 'self';
			frame-src 'none';
			base-uri 'self';
			form-action 'self';`, nonce),
		XFrameOptions:      "SAMEORIGIN",
		HSTSPreloadEnabled: false,
	})

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("nonce", nonce)

			return secureMiddlewareFunc(next)(c)
		}
	}
}

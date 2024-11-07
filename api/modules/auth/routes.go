package auth

import (
	"github.com/labstack/echo/v4"
)

func (m *AuthModule) Routes(g *echo.Group) {
	auth := g.Group("/auth")

	auth.GET("/register", m.registerHandler)
}

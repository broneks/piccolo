package auth

import (
	"github.com/labstack/echo/v4"
)

func (m *AuthModule) Routes(g *echo.Group) {
	auth := g.Group("/auth")

	auth.POST("/register", m.registerHandler)
	auth.POST("/login", m.loginHandler)
	auth.POST("/refresh", m.refreshHandler)
}

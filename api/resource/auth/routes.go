package auth

import (
	"github.com/labstack/echo/v4"
)

func (mod *AuthModule) Routes(g *echo.Group) {
	auth := g.Group("/auth")

	auth.POST("/register", mod.registerHandler)
	auth.POST("/login", mod.loginHandler)
	auth.POST("/refresh", mod.refreshHandler)
	auth.POST("/forgot-password", mod.forgotPasswordHandler)
}

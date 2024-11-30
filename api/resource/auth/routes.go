package auth

import (
	"os"

	"github.com/labstack/echo/v4"
)

func (mod *AuthModule) Routes(g *echo.Group) {
	auth := g.Group("/auth")

	if os.Getenv("ENV") != "local" {
		return
	}

	auth.POST("/register", mod.registerHandler)
	auth.POST("/login", mod.loginHandler)

	auth.POST("/refresh", mod.refreshHandler)
	auth.POST("/logout", mod.logoutHandler)
	auth.POST("/forgot-password", mod.forgotPasswordHandler)
	auth.PATCH("/reset-password", mod.resetPasswordHandler)
}

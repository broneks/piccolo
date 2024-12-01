package auth

import (
	"piccolo/api/middleware"

	"github.com/labstack/echo/v4"
)

func (mod *AuthModule) Routes(g *echo.Group) {
	auth := g.Group("/auth")

	auth.POST("/register", mod.registerHandler)
	auth.POST("/login", mod.loginHandler, middleware.BanHammer(mod.banHammerService))

	auth.POST("/refresh", mod.refreshHandler)
	auth.POST("/logout", mod.logoutHandler)
	auth.POST("/forgot-password", mod.forgotPasswordHandler)
	auth.PATCH("/reset-password", mod.resetPasswordHandler)
}

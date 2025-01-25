package me

import (
	"github.com/labstack/echo/v4"
)

func (mod *MeModule) Routes(g *echo.Group) {
	auth := g.Group("/me")

	auth.GET("/file-storage", mod.userFileStorageHandler)
}

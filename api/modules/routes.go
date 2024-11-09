package modules

import (
	"piccolo/api/middleware"
	"piccolo/api/modules/auth"
	"piccolo/api/modules/upload"
	"piccolo/api/repo"
	"piccolo/api/shared"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group, server *shared.Server) {
	v1 := g.Group("/v1")

	userRepo := repo.NewUserRepo(server.DB)
	photoRepo := repo.NewPhotoRepo(server.DB)

	// --- Public ---

	authModule := auth.New(server, userRepo)
	authModule.Routes(v1)

	// --- Protected ---

	photosGroup := v1.Group("/photos", middleware.Auth())

	uploadModule := upload.New(server, photoRepo)
	uploadModule.Routes(photosGroup)
}

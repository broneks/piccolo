package resource

import (
	"piccolo/api/middleware"
	"piccolo/api/repo"
	"piccolo/api/resource/album"
	"piccolo/api/resource/auth"
	"piccolo/api/resource/photo"
	"piccolo/api/service"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group, server *types.Server) {
	g.RouteNotFound("/*", handleNotFound)

	g.GET("/health", handleHealth(server))

	v1 := g.Group("/v1", middleware.SetUserData())

	userRepo := repo.NewUserRepo(server.DB)
	photoRepo := repo.NewPhotoRepo(server.DB)
	albumRepo := repo.NewAlbumRepo(server.DB)

	authService := service.NewAuthService(server, userRepo)
	photoService := service.NewPhotoService(server, photoRepo)

	authModule := auth.NewModule(server, userRepo, authService)
	authModule.Routes(v1)

	photoModule := photo.NewModule(server, photoRepo, photoService)
	photoModule.Routes(v1)

	albumModule := album.NewModule(server, albumRepo, photoService)
	albumModule.Routes(v1)
}

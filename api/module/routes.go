package module

import (
	"piccolo/api/middleware"
	"piccolo/api/module/album"
	"piccolo/api/module/auth"
	"piccolo/api/module/photo"
	"piccolo/api/repo"
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

	photoService := service.NewPhotoService(server, photoRepo)

	authModule := auth.New(server, userRepo)
	authModule.Routes(v1)

	photoModule := photo.New(server, photoRepo, photoService)
	photoModule.Routes(v1)

	albumModule := album.New(server, albumRepo, photoService)
	albumModule.Routes(v1)
}

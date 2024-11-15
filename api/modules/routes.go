package modules

import (
	"piccolo/api/modules/albums"
	"piccolo/api/modules/auth"
	"piccolo/api/modules/photos"
	photosservice "piccolo/api/modules/photos/service"
	"piccolo/api/repo"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group, server *types.Server) {
	g.GET("/health", handleHealth(server))

	v1 := g.Group("/v1")

	userRepo := repo.NewUserRepo(server.DB)
	photoRepo := repo.NewPhotoRepo(server.DB)
	albumRepo := repo.NewAlbumRepo(server.DB)

	photosService := photosservice.New(server, photoRepo)

	authModule := auth.New(server, userRepo)
	authModule.Routes(v1)

	photosModule := photos.New(server, photoRepo, photosService)
	photosModule.Routes(v1)

	albumsModule := albums.New(server, albumRepo, photosService)
	albumsModule.Routes(v1)
}

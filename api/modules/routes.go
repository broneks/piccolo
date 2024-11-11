package modules

import (
	"piccolo/api/modules/albums"
	"piccolo/api/modules/auth"
	"piccolo/api/modules/photos"
	"piccolo/api/modules/upload"
	"piccolo/api/repo"
	"piccolo/api/shared"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group, server *shared.Server) {
	v1 := g.Group("/v1")

	userRepo := repo.NewUserRepo(server.DB)
	photoRepo := repo.NewPhotoRepo(server.DB)
	albumRepo := repo.NewAlbumRepo(server.DB)

	authModule := auth.New(server, userRepo)
	authModule.Routes(v1)

	uploadModule := upload.New(server, photoRepo)
	uploadModule.Routes(v1)

	photosModule := photos.New(server, photoRepo)
	photosModule.Routes(v1)

	albumsModule := albums.New(server, albumRepo)
	albumsModule.Routes(v1)
}

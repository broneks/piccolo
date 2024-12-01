package resource

import (
	"piccolo/api/middleware"
	"piccolo/api/repo/albumrepo"
	"piccolo/api/repo/photorepo"
	"piccolo/api/repo/userrepo"
	"piccolo/api/resource/album"
	"piccolo/api/resource/auth"
	"piccolo/api/resource/photo"
	"piccolo/api/service/authservice"
	"piccolo/api/service/banhammerservice"
	"piccolo/api/service/photoservice"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group, server *types.Server) {
	g.GET("/health", handleHealth(server))

	v1 := g.Group("/v1", middleware.SetUserData())

	userRepo := userrepo.New(server.DB)
	photoRepo := photorepo.New(server.DB)
	albumRepo := albumrepo.New(server.DB)

	banHammerService := banhammerservice.New()
	authService := authservice.New(server, userRepo)
	photoService := photoservice.New(server, photoRepo)

	authModule := auth.NewModule(server, userRepo, banHammerService, authService)
	authModule.Routes(v1)

	photoModule := photo.NewModule(server, photoRepo, photoService)
	photoModule.Routes(v1)

	albumModule := album.NewModule(server, albumRepo, photoService)
	albumModule.Routes(v1)
}

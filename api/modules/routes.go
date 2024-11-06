package modules

import (
	"piccolo/api/modules/upload"
	"piccolo/api/repo"
	"piccolo/api/shared"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group, server *shared.Server) {
	v1 := g.Group("/v1")

	photoRepo := repo.NewPhotoRepo(server.DB)

	uploadModule := upload.New(server, photoRepo)
	uploadModule.Routes(v1)
}

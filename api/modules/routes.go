package modules

import (
	"piccolo/api/modules/upload"
	"piccolo/api/repo"
	"piccolo/api/shared"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, server *shared.Server) {
	photoRepo := repo.NewPhotoRepo(server.DB)

	uploadModule := upload.New(server, photoRepo)
	uploadModule.Routes(e)
}

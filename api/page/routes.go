package page

import (
	"piccolo/api/middleware"
	"piccolo/api/repo"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, server *types.Server) {
	sharedAlbumRepo := repo.NewSharedAlbumRepo(server.DB)

	e.GET("/reset-password", handleResetPasswordPage(), middleware.CanResetPassword())

	album := e.Group("/albums/:id", middleware.CanReadSharedAlbum(sharedAlbumRepo))

	album.GET("", handleSharedAlbumPage(server, sharedAlbumRepo))
}

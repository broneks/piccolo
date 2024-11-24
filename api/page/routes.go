package page

import (
	"piccolo/api/middleware"
	"piccolo/api/repo/sharedalbumrepo"
	"piccolo/api/repo/userrepo"
	"piccolo/api/service"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, server *types.Server) {
	userRepo := userrepo.New(server.DB)
	sharedAlbumRepo := sharedalbumrepo.New(server.DB)

	authService := service.NewAuthService(server, userRepo)

	resetPassword := e.Group("/reset-password", middleware.CanResetPassword(server))

	resetPassword.GET("", handleGetResetPasswordPage())
	resetPassword.POST("", handlePostResetPasswordPage(authService))

	album := e.Group("/albums/:id", middleware.CanReadSharedAlbum(sharedAlbumRepo))

	album.GET("", handleSharedAlbumPage(server, sharedAlbumRepo))
}

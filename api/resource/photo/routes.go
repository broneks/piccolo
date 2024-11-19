package photo

import (
	"piccolo/api/middleware"

	"github.com/labstack/echo/v4"
)

func (mod *PhotoModule) Routes(g *echo.Group) {
	photos := g.Group("/photos", middleware.Auth())

	photos.GET("", mod.getPhotosHandler)
	photos.POST("/upload", mod.postPhotosUploadHandler)

	photo := photos.Group("/:id")

	photo.GET("", mod.getPhotoHandler)
	photo.GET("/albums", mod.getPhotoAlbumsHandler)
}

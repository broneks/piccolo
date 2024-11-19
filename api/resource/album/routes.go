package album

import (
	"piccolo/api/middleware"

	"github.com/labstack/echo/v4"
)

func (mod *AlbumModule) Routes(g *echo.Group) {
	albums := g.Group("/albums", middleware.Auth())

	albums.GET("", mod.getAlbumsHandler)
	albums.POST("", mod.postAlbumsCreateHandler)

	album := albums.Group("/:id")

	album.GET("", mod.getAlbumHandler)
	album.GET("/users", mod.getAlbumUsersHandler)
	album.GET("/photos", mod.getAlbumPhotosHandler)
	album.POST("/upload", mod.postAlbumPhotosUploadHandler)
}

package albums

import (
	"piccolo/api/middleware"

	"github.com/labstack/echo/v4"
)

func (m *AlbumsModule) Routes(g *echo.Group) {
	albums := g.Group("/albums", middleware.Auth())

	albums.GET("", m.getAlbumsHandler)

	album := albums.Group("/:id")

	album.GET("", m.getAlbumHandler)
	album.GET("/users", m.getAlbumUsersHandler)
	album.GET("/photos", m.getAlbumPhotosHandler)
	album.POST("/upload", m.postAlbumPhotosUploadHandler)
}

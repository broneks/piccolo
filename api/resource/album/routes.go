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
	album.POST("/upload", mod.postAlbumPhotosUploadHandler)

	albumPhotos := album.Group("/photos")

	albumPhotos.GET("", mod.getAlbumPhotosHandler)
	albumPhotos.GET("/likes", mod.getAlbumPhotosLikesHandler)
	albumPhotos.GET("/favourites", mod.getAlbumPhotosFavouritesHandler)

	albumPhoto := albumPhotos.Group("/:photoId")

	albumPhoto.POST("/like", mod.postAlbumPhotoLike)
	albumPhoto.DELETE("/like", mod.deleteAlbumPhotoLike)

	albumPhoto.POST("/favourite", mod.postAlbumPhotoFavourite)
	albumPhoto.DELETE("/favourite", mod.deleteAlbumPhotoFavourite)
}

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
	albumPhotos.POST("", mod.postAlbumPhotoHandler)
	albumPhotos.GET("/likes", mod.getAlbumPhotosLikesHandler)
	albumPhotos.GET("/favourites", mod.getAlbumPhotosFavouritesHandler)

	albumPhoto := albumPhotos.Group("/:photoId")

	albumPhoto.DELETE("", mod.deleteAlbumPhotoHandler)

	albumPhoto.POST("/like", mod.postAlbumPhotoLikeHandler)
	albumPhoto.DELETE("/like", mod.deleteAlbumPhotoLikeHandler)

	albumPhoto.POST("/favourite", mod.postAlbumPhotoFavouriteHandler)
	albumPhoto.DELETE("/favourite", mod.deleteAlbumPhotoFavouriteHandler)
}

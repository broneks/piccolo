package albums

import (
	photos "piccolo/api/modules/photos/service"
	"piccolo/api/repo"
	"piccolo/api/types"
)

type AlbumsModule struct {
	server        *types.Server
	albumRepo     *repo.AlbumRepo
	photosService *photos.PhotosService
}

func New(server *types.Server, albumRepo *repo.AlbumRepo, photosService *photos.PhotosService) *AlbumsModule {
	return &AlbumsModule{
		server:        server,
		albumRepo:     albumRepo,
		photosService: photosService,
	}
}

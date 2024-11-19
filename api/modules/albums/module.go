package albums

import (
	"piccolo/api/repo"
	"piccolo/api/service"
	"piccolo/api/types"
)

type AlbumsModule struct {
	server       *types.Server
	albumRepo    *repo.AlbumRepo
	photoService *service.PhotoService
}

func New(server *types.Server, albumRepo *repo.AlbumRepo, photoService *service.PhotoService) *AlbumsModule {
	return &AlbumsModule{
		server:       server,
		albumRepo:    albumRepo,
		photoService: photoService,
	}
}

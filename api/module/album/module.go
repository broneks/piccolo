package album

import (
	"piccolo/api/repo"
	"piccolo/api/service"
	"piccolo/api/types"
)

type AlbumModule struct {
	server       *types.Server
	albumRepo    *repo.AlbumRepo
	photoService *service.PhotoService
}

func New(server *types.Server, albumRepo *repo.AlbumRepo, photoService *service.PhotoService) *AlbumModule {
	return &AlbumModule{
		server:       server,
		albumRepo:    albumRepo,
		photoService: photoService,
	}
}

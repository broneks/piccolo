package album

import (
	"piccolo/api/repo/albumrepo"
	"piccolo/api/service"
	"piccolo/api/types"
)

type AlbumModule struct {
	server       *types.Server
	albumRepo    *albumrepo.AlbumRepo
	photoService *service.PhotoService
}

func NewModule(server *types.Server, albumRepo *albumrepo.AlbumRepo, photoService *service.PhotoService) *AlbumModule {
	return &AlbumModule{
		server:       server,
		albumRepo:    albumRepo,
		photoService: photoService,
	}
}

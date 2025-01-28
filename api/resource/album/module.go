package album

import (
	"piccolo/api/repo/albumrepo"
	"piccolo/api/repo/photorepo"
	"piccolo/api/service/photoservice"
	"piccolo/api/types"
)

type AlbumModule struct {
	server       *types.Server
	albumRepo    *albumrepo.AlbumRepo
	photoRepo    *photorepo.PhotoRepo
	photoService *photoservice.PhotoService
}

func NewModule(server *types.Server, albumRepo *albumrepo.AlbumRepo, photoRepo *photorepo.PhotoRepo, photoService *photoservice.PhotoService) *AlbumModule {
	return &AlbumModule{
		server:       server,
		albumRepo:    albumRepo,
		photoRepo:    photoRepo,
		photoService: photoService,
	}
}

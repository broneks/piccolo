package photos

import (
	"piccolo/api/repo"
	"piccolo/api/service"
	"piccolo/api/types"
)

type PhotosModule struct {
	server       *types.Server
	photoRepo    *repo.PhotoRepo
	photoService *service.PhotoService
}

func New(server *types.Server, photoRepo *repo.PhotoRepo, photoService *service.PhotoService) *PhotosModule {
	return &PhotosModule{
		server:       server,
		photoRepo:    photoRepo,
		photoService: photoService,
	}
}

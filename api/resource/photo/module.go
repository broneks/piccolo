package photo

import (
	"piccolo/api/repo"
	"piccolo/api/service"
	"piccolo/api/types"
)

type PhotoModule struct {
	server       *types.Server
	photoRepo    *repo.PhotoRepo
	photoService *service.PhotoService
}

func NewModule(server *types.Server, photoRepo *repo.PhotoRepo, photoService *service.PhotoService) *PhotoModule {
	return &PhotoModule{
		server:       server,
		photoRepo:    photoRepo,
		photoService: photoService,
	}
}

package photo

import (
	"piccolo/api/repo/photorepo"
	"piccolo/api/service"
	"piccolo/api/types"
)

type PhotoModule struct {
	server       *types.Server
	photoRepo    *photorepo.PhotoRepo
	photoService *service.PhotoService
}

func NewModule(server *types.Server, photoRepo *photorepo.PhotoRepo, photoService *service.PhotoService) *PhotoModule {
	return &PhotoModule{
		server:       server,
		photoRepo:    photoRepo,
		photoService: photoService,
	}
}

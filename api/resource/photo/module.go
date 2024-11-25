package photo

import (
	"piccolo/api/repo/photorepo"
	"piccolo/api/service/photoservice"
	"piccolo/api/types"
)

type PhotoModule struct {
	server       *types.Server
	photoRepo    *photorepo.PhotoRepo
	photoService *photoservice.PhotoService
}

func NewModule(server *types.Server, photoRepo *photorepo.PhotoRepo, photoService *photoservice.PhotoService) *PhotoModule {
	return &PhotoModule{
		server:       server,
		photoRepo:    photoRepo,
		photoService: photoService,
	}
}

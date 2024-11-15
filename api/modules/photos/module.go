package photos

import (
	"piccolo/api/modules/photos/service"
	"piccolo/api/repo"
	"piccolo/api/types"
)

type PhotosModule struct {
	server        *types.Server
	photoRepo     *repo.PhotoRepo
	photosService *service.PhotosService
}

func New(server *types.Server, photoRepo *repo.PhotoRepo, photosService *service.PhotosService) *PhotosModule {
	return &PhotosModule{
		server:        server,
		photoRepo:     photoRepo,
		photosService: photosService,
	}
}

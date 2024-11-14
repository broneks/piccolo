package photos

import (
	"piccolo/api/repo"
	"piccolo/api/types"
)

type PhotosModule struct {
	server    *types.Server
	photoRepo *repo.PhotoRepo
}

func New(server *types.Server, photoRepo *repo.PhotoRepo) *PhotosModule {
	return &PhotosModule{
		server:    server,
		photoRepo: photoRepo,
	}
}

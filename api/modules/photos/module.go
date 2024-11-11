package photos

import (
	"piccolo/api/repo"
	"piccolo/api/shared"
)

type PhotosModule struct {
	server    *shared.Server
	photoRepo *repo.PhotoRepo
}

func New(server *shared.Server, photoRepo *repo.PhotoRepo) *PhotosModule {
	return &PhotosModule{
		server:    server,
		photoRepo: photoRepo,
	}
}

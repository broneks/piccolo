package upload

import (
	"piccolo/api/repo"
	"piccolo/api/shared"
)

type UploadModule struct {
	server    *shared.Server
	photoRepo *repo.PhotoRepo
}

func New(server *shared.Server, photoRepo *repo.PhotoRepo) *UploadModule {
	return &UploadModule{
		server:    server,
		photoRepo: photoRepo,
	}
}

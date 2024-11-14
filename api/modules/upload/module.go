package upload

import (
	"piccolo/api/repo"
	"piccolo/api/types"
)

type UploadModule struct {
	server    *types.Server
	photoRepo *repo.PhotoRepo
}

func New(server *types.Server, photoRepo *repo.PhotoRepo) *UploadModule {
	return &UploadModule{
		server:    server,
		photoRepo: photoRepo,
	}
}

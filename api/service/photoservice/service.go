package photoservice

import (
	"piccolo/api/repo/photorepo"
	"piccolo/api/types"
)

type PhotoService struct {
	server    *types.Server
	photoRepo *photorepo.PhotoRepo
}

func New(server *types.Server, photoRepo *photorepo.PhotoRepo) *PhotoService {
	return &PhotoService{
		server:    server,
		photoRepo: photoRepo,
	}
}


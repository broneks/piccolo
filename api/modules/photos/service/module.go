package service

import (
	"piccolo/api/repo"
	"piccolo/api/types"
)

type PhotosService struct {
	server    *types.Server
	photoRepo *repo.PhotoRepo
}

func New(server *types.Server, photoRepo *repo.PhotoRepo) *PhotosService {
	return &PhotosService{
		server:    server,
		photoRepo: photoRepo,
	}
}

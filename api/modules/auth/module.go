package auth

import (
	"piccolo/api/repo"
	"piccolo/api/shared"
)

type AuthModule struct {
	server    *shared.Server
	photoRepo *repo.PhotoRepo
}

func New(server *shared.Server, photoRepo *repo.PhotoRepo) *AuthModule {
	return &AuthModule{
		server:    server,
		photoRepo: photoRepo,
	}
}

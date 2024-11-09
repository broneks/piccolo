package auth

import (
	"piccolo/api/repo"
	"piccolo/api/shared"
)

type AuthModule struct {
	server   *shared.Server
	userRepo *repo.UserRepo
}

func New(server *shared.Server, userRepo *repo.UserRepo) *AuthModule {
	return &AuthModule{
		server:   server,
		userRepo: userRepo,
	}
}

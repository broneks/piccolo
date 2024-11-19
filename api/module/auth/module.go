package auth

import (
	"piccolo/api/repo"
	"piccolo/api/types"
)

type AuthModule struct {
	server   *types.Server
	userRepo *repo.UserRepo
}

func New(server *types.Server, userRepo *repo.UserRepo) *AuthModule {
	return &AuthModule{
		server:   server,
		userRepo: userRepo,
	}
}

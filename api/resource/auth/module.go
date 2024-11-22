package auth

import (
	"piccolo/api/repo"
	"piccolo/api/service"
	"piccolo/api/types"
)

type AuthModule struct {
	server      *types.Server
	userRepo    *repo.UserRepo
	authService *service.AuthService
}

func NewModule(server *types.Server, userRepo *repo.UserRepo, authService *service.AuthService) *AuthModule {
	return &AuthModule{
		server:      server,
		userRepo:    userRepo,
		authService: authService,
	}
}

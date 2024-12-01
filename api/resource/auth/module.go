package auth

import (
	"piccolo/api/repo/userrepo"
	"piccolo/api/service/authservice"
	"piccolo/api/service/banhammerservice"
	"piccolo/api/types"
)

type AuthModule struct {
	server           *types.Server
	userRepo         *userrepo.UserRepo
	banHammerService *banhammerservice.BanHammerService
	authService      *authservice.AuthService
}

func NewModule(
	server *types.Server,
	userRepo *userrepo.UserRepo,
	banHammerService *banhammerservice.BanHammerService,
	authService *authservice.AuthService,
) *AuthModule {
	return &AuthModule{
		server:           server,
		userRepo:         userRepo,
		banHammerService: banHammerService,
		authService:      authService,
	}
}

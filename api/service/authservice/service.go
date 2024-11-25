package authservice

import (
	"piccolo/api/repo/userrepo"
	"piccolo/api/types"
)

type AuthService struct {
	server                *types.Server
	userRepo              *userrepo.UserRepo
	MinPasswordCharLength int
}

func New(server *types.Server, userRepo *userrepo.UserRepo) *AuthService {
	return &AuthService{
		server:                server,
		userRepo:              userRepo,
		MinPasswordCharLength: 14,
	}
}

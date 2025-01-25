package me

import (
	"piccolo/api/service/photoservice"
	"piccolo/api/types"
)

type MeModule struct {
	server       *types.Server
	photoService *photoservice.PhotoService
}

func NewModule(server *types.Server, photoService *photoservice.PhotoService) *MeModule {
	return &MeModule{
		server:       server,
		photoService: photoService,
	}
}

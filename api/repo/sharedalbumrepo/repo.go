package sharedalbumrepo

import (
	"piccolo/api/types"
)

type SharedAlbumRepo struct {
	db types.ServerDB
}

func New(db types.ServerDB) *SharedAlbumRepo {
	return &SharedAlbumRepo{db: db}
}

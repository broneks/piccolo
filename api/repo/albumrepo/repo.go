package albumrepo

import (
	"piccolo/api/types"
)

type AlbumRepo struct {
	db types.ServerDB
}

func New(db types.ServerDB) *AlbumRepo {
	return &AlbumRepo{db: db}
}

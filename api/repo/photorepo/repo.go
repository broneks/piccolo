package photorepo

import (
	"piccolo/api/types"
)

type PhotoRepo struct {
	db types.ServerDB
}

func New(db types.ServerDB) *PhotoRepo {
	return &PhotoRepo{db: db}
}

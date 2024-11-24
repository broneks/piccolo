package userrepo

import (
	"piccolo/api/types"
)

type UserRepo struct {
	db types.ServerDB
}

func New(db types.ServerDB) *UserRepo {
	return &UserRepo{db: db}
}

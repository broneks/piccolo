package repo

import "piccolo/api/shared"

type UserRepo struct {
	db shared.ServerDB
}

func NewUserRepo(db shared.ServerDB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) GetById(id string) (string, error) {
	return "", nil
}

func (r *UserRepo) GetByEmail(email string) (string, error) {
	return "", nil
}

func (r *UserRepo) InsertOne(email, password string) error {
	return nil
}

func (r *UserRepo) UpdateLastLoginAt(id string) error {
	return nil
}

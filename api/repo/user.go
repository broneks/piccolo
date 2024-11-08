package repo

import (
	"context"
	"fmt"
	"piccolo/api/model"
	"piccolo/api/shared"

	"github.com/jackc/pgx/v5"
)

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

func (r *UserRepo) InsertOne(ctx context.Context, user model.User) error {
	query := `insert into users (
		username,
		email,
		hash,
		hashedAt,
	) values (
		@username,
		@email,
		@hash,
		@hashedAt,
	)`

	args := pgx.NamedArgs{
		"username": user.Username,
		"email":    user.Email,
		"hash":     user.Hash,
		"hashedAt": user.HashedAt,
	}
	_, err := r.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (r *UserRepo) UpdateLastLoginAt(id string) error {
	return nil
}

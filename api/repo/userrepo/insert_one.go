package userrepo

import (
	"context"
	"fmt"
	"piccolo/api/model"

	"github.com/jackc/pgx/v5"
)

func (repo *UserRepo) InsertOne(ctx context.Context, user model.User) error {
	query := `insert into users (
		username,
		email,
		hash,
		hashed_at
	) values (
		@username,
		@email,
		@hash,
		@hashedAt
	)`

	args := pgx.NamedArgs{
		"username": user.Username,
		"email":    user.Email,
		"hash":     user.Hash,
		"hashedAt": user.HashedAt,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

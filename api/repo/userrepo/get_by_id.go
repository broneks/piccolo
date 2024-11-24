package userrepo

import (
	"context"
	"fmt"
	"piccolo/api/model"

	"github.com/jackc/pgx/v5"
)

func (repo *UserRepo) GetById(ctx context.Context, id string) (*model.User, error) {
	query := `select
		id,
		username,
		email,
		hash,
		hashed_at,
		last_login_at,
		created_at,
		updated_at
	from users where id = $1`

	var user model.User

	err := repo.db.QueryRow(ctx, query, id).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Hash,
		&user.HashedAt,
		&user.LastLoginAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user with id %s not found", id)
		}
		return nil, fmt.Errorf("query error: %v", err)
	}

	return &user, nil
}

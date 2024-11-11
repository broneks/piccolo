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

func (r *UserRepo) GetById(id string) (*model.User, error) {
	query := `select * from users where id = $1`

	var user model.User

	err := r.db.QueryRow(context.Background(), query, id).Scan(
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

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `select * from users where email = $1`

	var user model.User

	err := r.db.QueryRow(context.Background(), query, email).Scan(
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
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, fmt.Errorf("query error: %v", err)
	}

	return &user, nil
}

func (r *UserRepo) InsertOne(ctx context.Context, user model.User) error {
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
	_, err := r.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (r *UserRepo) UpdateLastLoginAt(ctx context.Context, userId string) error {
	query := `update users set last_login_at = now() where id = $1`

	_, err := r.db.Exec(ctx, query, userId)
	if err != nil {
		return fmt.Errorf("unable to update row: %w", err)
	}

	return nil
}

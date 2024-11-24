package userrepo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (repo *UserRepo) UpdatePassword(ctx context.Context, userId, hash string) error {
	query := `update users set hash = @hash, hashed_at = now() where id = @userId`

	args := pgx.NamedArgs{
		"hash":   hash,
		"userId": userId,
	}
	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

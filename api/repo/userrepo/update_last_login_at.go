package userrepo

import (
	"context"
	"fmt"
)

func (repo *UserRepo) UpdateLastLoginAt(ctx context.Context, userId string) error {
	query := `update users set last_login_at = now() where id = $1`

	_, err := repo.db.Exec(ctx, query, userId)
	if err != nil {
		return fmt.Errorf("unable to update row: %w", err)
	}

	return nil
}

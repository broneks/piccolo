package photorepo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// remove photos uploaded by user
func (repo *PhotoRepo) RemoveMany(ctx context.Context, photoIds []string, userId string) (int64, error) {
	query := `delete from photos where id = any(@photoIds) and user_id = @userId`

	args := pgx.NamedArgs{
		"photoIds": photoIds,
		"userId":   userId,
	}

	cmd, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return 0, fmt.Errorf("unable to delete rows: %w", err)
	}

	return cmd.RowsAffected(), nil
}

package photorepo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// remove photo uploaded by user
func (repo *PhotoRepo) RemoveOne(ctx context.Context, photoId, userId string) (int64, error) {
	query := `delete from photos where id = @photoId and user_id = @userId`

	args := pgx.NamedArgs{
		"photoId": photoId,
		"userId":  userId,
	}

	cmd, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return 0, fmt.Errorf("unable to delete photo: %w", err)
	}

	return cmd.RowsAffected(), nil
}

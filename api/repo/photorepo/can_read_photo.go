package photorepo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Permits:
// - photo uploader
func (repo *PhotoRepo) CanReadPhoto(ctx context.Context, photoId, userId string) (bool, error) {
	query := `select exists (
		select 1 from photos where id = @photoId and user_id = @userId
	) as can`

	var can bool

	args := pgx.NamedArgs{
		"photoId": photoId,
		"userId":  userId,
	}
	err := repo.db.QueryRow(ctx, query, args).Scan(&can)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, fmt.Errorf("no rows found for photo id '%s' and user id '%s'", photoId, userId)
		}
		return false, fmt.Errorf("query error: %v", err)
	}

	return can, nil
}

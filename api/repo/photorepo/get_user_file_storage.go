package photorepo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Get the total used file storage in MB for photos uploaded by the user
func (repo *PhotoRepo) GetUserFileStorage(ctx context.Context, userId string) (float32, error) {
	query := `select
		cast(sum(photos.file_size) as DECIMAL(10,2)) / (1024 * 1024) as used_mb
	from photos where user_id = @userId`

	var usedMB *float32

	args := pgx.NamedArgs{
		"userId": userId,
	}
	err := repo.db.QueryRow(ctx, query, args).Scan(&usedMB)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, fmt.Errorf("no rows found for user id '%s'", userId)
		}
		return 0, fmt.Errorf("query error: %v", err)
	}

	if usedMB == nil {
		return 0, nil
	}

	return *usedMB, nil
}

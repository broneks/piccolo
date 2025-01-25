package photorepo

import (
	"context"
	"fmt"
	"piccolo/api/model"

	"github.com/jackc/pgx/v5"
)

// Get photos uploaded by the user
func (repo *PhotoRepo) GetByIds(ctx context.Context, photoIds []string, userId string) ([]model.Photo, error) {
	var err error

	query := `select
		id,
		user_id,
		location,
		filename,
		file_size,
		content_type,
		created_at,
		updated_at
	from photos where id = any(@photoIds) and user_id = @userId`

	args := pgx.NamedArgs{
		"photoIds": photoIds,
		"userId":   userId,
	}

	rows, err := repo.db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to query photos: %v", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Photo])
}

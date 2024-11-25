package photorepo

import (
	"context"
	"fmt"
	"piccolo/api/model"

	"github.com/jackc/pgx/v5"
)

// Get all photos uploaded by the user
func (repo *PhotoRepo) GetAll(ctx context.Context, userId string) ([]model.Photo, error) {
	query := `select
		id,
		user_id,
		location,
		filename,
		file_size,
		content_type,
		created_at,
		updated_at
	from photos where user_id = $1
	order by created_at desc`

	rows, err := repo.db.Query(ctx, query, userId)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %v", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Photo])
}

package photorepo

import (
	"context"
	"fmt"
	"piccolo/api/model"
	"piccolo/api/types"

	"github.com/jackc/pgx/v5"
)

// Get all photos uploaded by the user
func (repo *PhotoRepo) GetAll(ctx context.Context, userId string) ([]model.Photo, error) {
	return repo.GetAllWithParams(ctx, userId, types.NewDefaultListQueryParams())
}

func (repo *PhotoRepo) GetAllWithParams(ctx context.Context, userId string, queryParams types.ListQueryParams) ([]model.Photo, error) {
	query := queryParams.WrapQuery(`select
		id,
		user_id,
		location,
		filename,
		file_size,
		content_type,
		created_at,
		updated_at
	from photos where user_id = @userId
	order by created_at desc`)

	args := queryParams.WrapNamedArgs(pgx.NamedArgs{
		"userId": userId,
	})

	rows, err := repo.db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %v", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Photo])
}

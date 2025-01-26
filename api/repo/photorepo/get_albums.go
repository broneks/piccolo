package photorepo

import (
	"context"
	"fmt"
	"piccolo/api/model"
	"piccolo/api/types"

	"github.com/jackc/pgx/v5"
)

func (repo *PhotoRepo) GetAlbums(ctx context.Context, photoId, userId string) ([]model.Album, error) {
	return repo.GetAlbumsWithParams(ctx, photoId, userId, types.NewDefaultListQueryParams())
}

func (repo *PhotoRepo) GetAlbumsWithParams(ctx context.Context, photoId, userId string, queryParams types.ListQueryParams) ([]model.Album, error) {
	query := queryParams.WrapQuery(`select
		a.id,
		a.user_id,
		a.name,
		a.description,
		a.cover_photo_id,
		a.is_share_link_enabled,
		a.read_access_hash,
		a.created_at,
		a.updated_at
	from albums a
	join album_photos ap on a.id = ap.album_id
	where ap.photo_id = @photoId
	and ap.user_id = @userId
	order by a.created_at desc`)

	args := queryParams.WrapNamedArgs(pgx.NamedArgs{
		"photoId": photoId,
		"userId":  userId,
	})

	rows, err := repo.db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to query albums: %v", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Album])
}

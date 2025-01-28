package albumrepo

import (
	"context"
	"fmt"
	"log/slog"
	"piccolo/api/model"
	"piccolo/api/types"

	"github.com/jackc/pgx/v5"
)

func (repo *AlbumRepo) GetPhotos(ctx context.Context, albumId, userId string) ([]model.AlbumPhoto, error) {
	return repo.GetPhotosWithParams(ctx, albumId, userId, types.NewDefaultListQueryParams())
}

// Checks for read access
func (repo *AlbumRepo) GetPhotosWithParams(ctx context.Context, albumId, userId string, queryParams types.ListQueryParams) ([]model.AlbumPhoto, error) {
	var err error

	canRead, err := repo.CanReadAlbum(ctx, albumId, userId)
	if err != nil {
		return nil, err
	}
	if !canRead {
		return nil, fmt.Errorf("unauthorized")
	}

	query := queryParams.WrapQuery(`select
		p.id,
		p.user_id,
		p.location,
		p.filename,
		p.file_size,
		p.content_type,
		p.created_at,
		p.updated_at,
		ap.created_at as added_at
	from photos p
	join album_photos ap on p.id = ap.photo_id
	where ap.album_id = @albumId
	order by p.created_at desc`)

	args := queryParams.WrapNamedArgs(pgx.NamedArgs{
		"albumId": albumId,
	})

	rows, err := repo.db.Query(ctx, query, args)
	if err != nil {
		slog.Debug(err.Error())
		return nil, fmt.Errorf("unable to query photos: %v", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.AlbumPhoto])
}

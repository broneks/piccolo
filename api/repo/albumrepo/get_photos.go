package albumrepo

import (
	"context"
	"fmt"
	"log/slog"
	"piccolo/api/model"

	"github.com/jackc/pgx/v5"
)

// Checks for read access
func (repo *AlbumRepo) GetPhotos(ctx context.Context, albumId, userId string) ([]model.Photo, error) {
	var err error

	canRead, err := repo.CanReadAlbum(ctx, albumId, userId)
	if err != nil {
		return nil, err
	}
	if !canRead {
		return nil, fmt.Errorf("unauthorized")
	}

	query := `select
		p.id,
		p.user_id,
		p.location,
		p.filename,
		p.file_size,
		p.content_type,
		p.created_at,
		p.updated_at
	from photos p
	join album_photos ap on p.id = ap.photo_id
	where ap.album_id = $1
	order by p.created_at desc`

	rows, err := repo.db.Query(ctx, query, albumId)
	if err != nil {
		slog.Debug(err.Error())
		return nil, fmt.Errorf("unable to query photos: %v", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Photo])
}

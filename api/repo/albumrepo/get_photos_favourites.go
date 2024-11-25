package albumrepo

import (
	"context"
	"fmt"
	"log/slog"
	"piccolo/api/model"

	"github.com/jackc/pgx/v5"
)

// Checks for read access
func (repo *AlbumRepo) GetPhotosFavourites(ctx context.Context, albumId, userId string) ([]model.AlbumPhotoFavourites, error) {
	var err error

	canRead, err := repo.CanReadAlbum(ctx, albumId, userId)
	if err != nil {
		return nil, err
	}
	if !canRead {
		return nil, fmt.Errorf("unauthorized")
	}

	query := `select photo_id, created_at
		from album_photo_favourites
		where album_id = @albumId and user_id = @userId`

	args := pgx.NamedArgs{
		"albumId": albumId,
		"userId":  userId,
	}

	rows, err := repo.db.Query(ctx, query, args)
	if err != nil {
		slog.Debug(err.Error())
		return nil, fmt.Errorf("unable to query album photo favourites: %v", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.AlbumPhotoFavourites])
}

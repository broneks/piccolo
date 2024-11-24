package albumrepo

import (
	"context"
	"fmt"
	"log/slog"
	"piccolo/api/model"

	"github.com/jackc/pgx/v5"
)

// Checks for read access
func (repo *AlbumRepo) GetPhotosLikes(ctx context.Context, albumId string, photoIds []string, userId string) ([]model.AlbumPhotoLikes, error) {
	var err error

	canRead, err := repo.CanReadAlbum(ctx, albumId, userId)
	if err != nil {
		return nil, err
	}
	if !canRead {
		return nil, fmt.Errorf("unauthorized")
	}

	var query string

	if len(photoIds) > 0 {
		query = `select photo_id, count(*) as likes
			from album_photo_likes
			where album_id = @albumId and photo_id = any(@photoIds)
			group by photo_id`
	} else {
		query = `select photo_id, count(*) as likes
			from album_photo_likes
			where album_id = @albumId group by photo_id`
	}

	args := pgx.NamedArgs{
		"albumId":  albumId,
		"photoIds": photoIds,
	}

	rows, err := repo.db.Query(ctx, query, args)
	if err != nil {
		slog.Debug(err.Error())
		return nil, fmt.Errorf("unable to query album photo likes: %v", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.AlbumPhotoLikes])
}

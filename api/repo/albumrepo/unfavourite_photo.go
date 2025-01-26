package albumrepo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Checks for read access
func (repo *AlbumRepo) UnfavouritePhoto(ctx context.Context, albumId, photoId, userId string) (int64, error) {

	var err error

	canRead, err := repo.CanReadAlbum(ctx, albumId, userId)
	if err != nil {
		return 0, err
	}
	if !canRead {
		return 0, fmt.Errorf("unauthorized")
	}

	query := `delete from album_photo_favourites
		where album_id = @albumId
		and photo_id = @photoId
		and user_id = @userId`

	args := pgx.NamedArgs{
		"albumId": albumId,
		"photoId": photoId,
		"userId":  userId,
	}

	cmd, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return 0, fmt.Errorf("unable to delete row: %w", err)
	}

	return cmd.RowsAffected(), nil
}

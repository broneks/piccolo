package albumrepo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Checks for read access
func (repo *AlbumRepo) UnlikePhoto(ctx context.Context, albumId, photoId, userId string) error {

	var err error

	canRead, err := repo.CanReadAlbum(ctx, albumId, userId)
	if err != nil {
		return err
	}
	if !canRead {
		return fmt.Errorf("unauthorized")
	}

	query := `delete from album_photo_likes
		where album_id = @albumId
		and photo_id = @photoId
		and user_id = @userId`

	args := pgx.NamedArgs{
		"albumId": albumId,
		"photoId": photoId,
		"userId":  userId,
	}

	_, err = repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to delete row: %w", err)
	}

	return nil
}

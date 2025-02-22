package albumrepo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Checks for read access
func (repo *AlbumRepo) FavouritePhoto(ctx context.Context, albumId, photoId, userId string) error {
	var err error

	canRead, err := repo.CanReadAlbum(ctx, albumId, userId)
	if err != nil {
		return err
	}
	if !canRead {
		return fmt.Errorf("unauthorized")
	}

	query := `insert into album_photo_favourites (
		album_id,
		photo_id,
		user_id
	) values (
		@albumId,
		@photoId,
		@userId
	)`

	args := pgx.NamedArgs{
		"albumId": albumId,
		"photoId": photoId,
		"userId":  userId,
	}

	_, err = repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

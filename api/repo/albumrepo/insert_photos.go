package albumrepo

import (
	"context"
	"fmt"
	"log/slog"
	"piccolo/api/helper"

	"github.com/jackc/pgx/v5"
)

// Checks for write access
func (repo *AlbumRepo) InsertPhotos(ctx context.Context, albumId string, photoIds []string, userId string) error {
	var err error

	canWrite, err := repo.CanWriteAlbum(ctx, albumId, userId)
	if err != nil {
		return err
	}
	if !canWrite {
		return fmt.Errorf("unauthorized")
	}

	query := `insert into album_photos (
		album_id,
		photo_id,
		user_id
	) values (
		@albumId,
		@photoId,
		@userId
	)`

	batch := &pgx.Batch{}

	for _, photoId := range photoIds {
		args := pgx.NamedArgs{
			"albumId": albumId,
			"photoId": photoId,
			"userId":  userId,
		}
		batch.Queue(query, args)
	}

	results := repo.db.SendBatch(ctx, batch)
	defer results.Close()

	for _, photoId := range photoIds {
		_, err := results.Exec()
		if err != nil {
			if sqlErr := helper.CheckSqlError(err); sqlErr == "unique-violation" {
				slog.Debug(fmt.Sprintf("photo \"%s\" already exists", photoId))

				if len(photoIds) == 1 {
					return err
				} else {
					continue
				}
			}

			slog.Debug(err.Error())
			return fmt.Errorf("unable to insert album photo: %w", err)
		}
	}

	return nil
}

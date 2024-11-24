package albumrepo

import (
	"context"
	"fmt"
	"log/slog"
	"piccolo/api/helper"
	"piccolo/api/model"

	"github.com/jackc/pgx/v5"
)

// Checks for write access
func (repo *AlbumRepo) InsertUsers(ctx context.Context, albumId string, albumUsers []model.AlbumUser, userId string) error {
	var err error

	canWrite, err := repo.CanWriteAlbum(ctx, albumId, userId)
	if err != nil {
		return err
	}
	if !canWrite {
		return fmt.Errorf("unauthorized")
	}

	query := `insert into album_users (
		album_id,
		user_id,
		role
	) values (
		@albumId,
		@userId,
		@role
	)`

	batch := &pgx.Batch{}

	for _, albumUser := range albumUsers {
		args := pgx.NamedArgs{
			"albumId": albumId,
			"userId":  albumUser.UserId,
			"role":    albumUser.Role,
		}
		batch.Queue(query, args)
	}

	results := repo.db.SendBatch(ctx, batch)
	defer results.Close()

	for _, albumUser := range albumUsers {
		_, err := results.Exec()
		if err != nil {
			if sqlErr := helper.CheckSqlError(err); sqlErr == "unique-violation" {
				slog.Debug(fmt.Sprintf("album user \"%s\" already exists", albumUser.UserId.String))
				continue
			}

			slog.Debug(err.Error())
			return fmt.Errorf("unable to insert album user: %w", err)
		}
	}

	return nil
}

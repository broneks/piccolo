package albumrepo

import (
	"context"
	"fmt"
	"log/slog"
	"piccolo/api/model"

	"github.com/jackc/pgx/v5"
)

// Permits all users
func (repo *AlbumRepo) InsertOne(ctx context.Context, album model.Album) error {
	query := `insert into albums (
		user_id,
		name,
		description,
		cover_photo_id,
		is_share_link_enabled,
		read_access_hash
	) values (
		@userId,
		@name,
		@description,
		@coverPhotoId,
		@isShareLinkEnabled,
		@readAccessHash
	)`

	args := pgx.NamedArgs{
		"userId":             album.UserId,
		"name":               album.Name,
		"description":        album.Description,
		"coverPhotoId":       album.CoverPhotoId,
		"isShareLinkEnabled": album.IsShareLinkEnabled,
		"readAccessHash":     album.ReadAccessHash,
	}
	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		slog.Debug(err.Error())
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

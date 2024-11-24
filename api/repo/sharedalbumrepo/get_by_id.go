package sharedalbumrepo

import (
	"context"
	"fmt"
	"piccolo/api/model"

	"github.com/jackc/pgx/v5"
)

func (repo *SharedAlbumRepo) GetById(ctx context.Context, albumId string) (*model.Album, error) {
	query := `select
		id,
		user_id,
		name,
		description,
		cover_photo_id,
		created_at,
		updated_at
	from albums where id = $1`

	var err error
	var album model.Album

	err = repo.db.QueryRow(ctx, query, albumId).Scan(
		&album.Id,
		&album.UserId,
		&album.Name,
		&album.Description,
		&album.CoverPhotoId,
		&album.CreatedAt,
		&album.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("album with id %s not found", albumId)
		}
		return nil, fmt.Errorf("query error: %v", err)
	}

	return &album, nil
}

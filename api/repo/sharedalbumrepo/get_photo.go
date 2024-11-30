package sharedalbumrepo

import (
	"context"
	"fmt"
	"piccolo/api/model"

	"github.com/jackc/pgx/v5"
)

func (repo *SharedAlbumRepo) GetPhoto(ctx context.Context, albumId, photoId string) (*model.Photo, error) {
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
	where ap.album_id = @albumId
	and ap.photo_id = @photoId`

	var err error
	var photo model.Photo

	args := pgx.NamedArgs{
		"albumId": albumId,
		"photoId": photoId,
	}

	err = repo.db.QueryRow(ctx, query, args).Scan(
		&photo.Id,
		&photo.UserId,
		&photo.Location,
		&photo.Filename,
		&photo.FileSize,
		&photo.ContentType,
		&photo.CreatedAt,
		&photo.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("photo with id %s not found", photoId)
		}
		return nil, fmt.Errorf("query error: %v", err)
	}

	return &photo, nil
}
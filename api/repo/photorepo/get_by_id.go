package photorepo

import (
	"context"
	"fmt"
	"piccolo/api/model"

	"github.com/jackc/pgx/v5"
)

// Get photo uploaded by the user
func (repo *PhotoRepo) GetById(ctx context.Context, photoId, userId string) (*model.Photo, error) {
	var err error

	query := `select
		id,
		user_id,
		location,
		filename,
		file_size,
		content_type,
		created_at,
		updated_at
	from photos where id = @photoId and user_id = @userId`

	var photo model.Photo

	args := pgx.NamedArgs{
		"photoId": photoId,
		"userId":  userId,
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

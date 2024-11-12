package repo

import (
	"context"
	"fmt"
	"piccolo/api/model"
	"piccolo/api/shared"

	"github.com/jackc/pgx/v5"
)

type PhotoRepo struct {
	db shared.ServerDB
}

func NewPhotoRepo(db shared.ServerDB) *PhotoRepo {
	return &PhotoRepo{db: db}
}

// Get all photo for which the user is the uploader
func (r *PhotoRepo) GetAll(ctx context.Context, userId string) ([]model.Photo, error) {
	query := `select * from photos where user_id = $1`

	rows, err := r.db.Query(ctx, query, userId)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %v", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Photo])
}

func (r *PhotoRepo) GetAlbums(ctx context.Context, photoId, userId string) ([]model.Album, error) {
	query := `select a.id, a.user_id, a.name, a.description, a.cover_photo_id, a.read_access_hash, a.created_at, a.updated_at
						from albums a
						join album_photos ap on p.id = ap.album_id
						where ap.photo_id = @photoId
						and ap.user_id = @userId`

	args := pgx.NamedArgs{
		"photoId": photoId,
		"userId":  userId,
	}
	rows, err := r.db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to query albums: %v", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Album])
}

func (r *PhotoRepo) InsertOne(ctx context.Context, photo model.Photo) error {
	query := `insert into photos (
		user_id,
		location,
		filename,
		file_size,
		content_type
	) values (
		@userId,
		@location,
		@filename,
		@fileSize,
		@contentType
	)`

	args := pgx.NamedArgs{
		"userId":      photo.UserId,
		"location":    photo.Location,
		"filename":    photo.Filename,
		"fileSize":    photo.FileSize,
		"contentType": photo.ContentType,
	}
	_, err := r.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

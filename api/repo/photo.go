package repo

import (
	"context"
	"fmt"
	"piccolo/api/model"
	"piccolo/api/types"

	"github.com/jackc/pgx/v5"
)

type PhotoRepo struct {
	db types.ServerDB
}

func NewPhotoRepo(db types.ServerDB) *PhotoRepo {
	return &PhotoRepo{db: db}
}

// Get photo uploaded by the user
func (r *PhotoRepo) GetById(ctx context.Context, photoId, userId string) (*model.Photo, error) {
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
	err = r.db.QueryRow(ctx, query, args).Scan(
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

// Get all photos uploaded by the user
func (r *PhotoRepo) GetAll(ctx context.Context, userId string) ([]model.Photo, error) {
	query := `select
		id,
		user_id,
		location,
		filename,
		file_size,
		content_type,
		created_at,
		updated_at
	from photos where user_id = $1`

	rows, err := r.db.Query(ctx, query, userId)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %v", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Photo])
}

func (r *PhotoRepo) GetAlbums(ctx context.Context, photoId, userId string) ([]model.Album, error) {
	query := `select
		a.id,
		a.user_id,
		a.name,
		a.description,
		a.cover_photo_id,
		a.read_access_hash,
		a.created_at,
		a.updated_at
	from albums a
	join album_photos ap on a.id = ap.album_id
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

func (r *PhotoRepo) InsertMany(ctx context.Context, photos []model.Photo) error {
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

	batch := &pgx.Batch{}

	for _, photo := range photos {
		args := pgx.NamedArgs{
			"userId":      photo.UserId,
			"location":    photo.Location,
			"filename":    photo.Filename,
			"fileSize":    photo.FileSize,
			"contentType": photo.ContentType,
		}
		batch.Queue(query, args)
	}

	results := r.db.SendBatch(ctx, batch)
	defer results.Close()

	for _, photo := range photos {
		_, err := results.Exec()
		if err != nil {
			return fmt.Errorf("unable to insert photo \"%s\": %w", photo.Filename.String, err)
		}
	}

	return nil
}

// TODO
func (r *PhotoRepo) Update(ctx context.Context, photo model.Photo, userId string) error {
	return nil
}

// TODO
// remove photo uploaded by user
func (r *PhotoRepo) RemoveOne(ctx context.Context, photoId, userId string) error {
	return nil
}

// TODO
// remove photos uploaded by user
func (r *PhotoRepo) RemoveMany(ctx context.Context, photoIds []string, userId string) error {
	return nil
}

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

func (r *PhotoRepo) GetAll(ctx context.Context) ([]model.Photo, error) {
	query := `select * from photos`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %v", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Photo])
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

package pg

import (
	"context"
	"fmt"
	"piccolo/api/model"

	"github.com/jackc/pgx/v5"
)

func (pg *postgres) GetPhotos(ctx context.Context) ([]model.Photo, error) {
	query := `select * from photos`

	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %v", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Photo])
}

func (pg *postgres) InsertPhoto(ctx context.Context, photo model.Photo) error {
	query := `insert into photos (
		location,
		filename,
		file_size,
		content_type
	) values (
		@location,
		@filename,
		@fileSize,
		@contentType
	)`

	args := pgx.NamedArgs{
		"location":    photo.Location,
		"filename":    photo.Filename,
		"fileSize":    photo.FileSize,
		"contentType": photo.ContentType,
	}
	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

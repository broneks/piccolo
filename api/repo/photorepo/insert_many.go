package photorepo

import (
	"context"
	"fmt"
	"piccolo/api/model"

	"github.com/jackc/pgx/v5"
)

func (repo *PhotoRepo) InsertMany(ctx context.Context, photos []model.Photo, userId string) ([]string, error) {
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
	) returning id`

	batch := &pgx.Batch{}

	for _, photo := range photos {
		args := pgx.NamedArgs{
			"userId":      userId,
			"location":    photo.Location,
			"filename":    photo.Filename,
			"fileSize":    photo.FileSize,
			"contentType": photo.ContentType,
		}
		batch.Queue(query, args)
	}

	results := repo.db.SendBatch(ctx, batch)
	defer results.Close()

	var ids []string

	for _, photo := range photos {
		var id string

		if err := results.QueryRow().Scan(&id); err != nil {
			return nil, fmt.Errorf("unable to fetch inserted id for photo \"%s\": %w", photo.Filename.String, err)
		}

		ids = append(ids, id)
	}

	return ids, nil
}

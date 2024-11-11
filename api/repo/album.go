package repo

import (
	"context"
	"fmt"
	"piccolo/api/model"
	"piccolo/api/shared"

	"github.com/jackc/pgx/v5"
)

type AlbumRepo struct {
	db shared.ServerDB
}

func NewAlbumRepo(db shared.ServerDB) *AlbumRepo {
	return &AlbumRepo{db: db}
}

func (r *AlbumRepo) GetAll(ctx context.Context, userId string) ([]model.Album, error) {
	query := `select * from albums where user_id = $1`

	rows, err := r.db.Query(ctx, query, userId)
	if err != nil {
		return nil, fmt.Errorf("unable to query albums: %v", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Album])
}

func (r *AlbumRepo) GetPhotos(ctx context.Context, albumId string) ([]model.Photo, error) {
	query := `select p.id, p.user_id, p.location, p.filename, p.file_size, p.content_type, p.created_at, p.updated_at
						from photos p
						join album_photos ap on p.id = ap.photo_id
						where ap.album_id = $1`

	rows, err := r.db.Query(ctx, query, albumId)
	if err != nil {
		return nil, fmt.Errorf("unable to query photos: %v", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Photo])
}

func (r *AlbumRepo) InsertOne(ctx context.Context, album model.Album) error {
	query := `insert into albums (
		user_id,
		name,
		description,
		cover_photo_id
	) values (
		@userId,
		@name,
		@description,
		@coverPhotoId
	)`

	args := pgx.NamedArgs{
		"userId":       album.UserId,
		"name":         album.Name,
		"description":  album.Description,
		"coverPhotoId": album.CoverPhotoId,
	}
	_, err := r.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

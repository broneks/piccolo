package repo

import (
	"context"
	"fmt"
	"log/slog"
	"piccolo/api/model"
	"piccolo/api/types"

	"github.com/jackc/pgx/v5"
)

type SharedAlbumRepo struct {
	db types.ServerDB
}

func NewSharedAlbumRepo(db types.ServerDB) *SharedAlbumRepo {
	return &SharedAlbumRepo{db: db}
}

// Permits:
// - non-users with a valid album read access hash for an album that has sharing enabled
func (r *SharedAlbumRepo) CanReadSharedAlbum(ctx context.Context, albumId, readAccessHash string) (bool, error) {
	query := `select exists (
		select 1 from albums
		where id = @albumId and is_share_link_enabled = true and read_access_hash = @readAccessHash
	) as can`

	var can bool

	args := pgx.NamedArgs{
		"albumId":        albumId,
		"readAccessHash": readAccessHash,
	}

	err := r.db.QueryRow(ctx, query, args).Scan(&can)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, fmt.Errorf("no rows found for album id '%s'", albumId)
		}
		return false, fmt.Errorf("query error: %v", err)
	}

	return can, nil
}

func (r *SharedAlbumRepo) GetById(ctx context.Context, albumId string) (*model.Album, error) {
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

	err = r.db.QueryRow(ctx, query, albumId).Scan(
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

func (r *SharedAlbumRepo) GetPhotos(ctx context.Context, albumId string) ([]model.Photo, error) {
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
	where ap.album_id = $1
	order by p.created_at desc`

	rows, err := r.db.Query(ctx, query, albumId)
	if err != nil {
		slog.Debug(err.Error())
		return nil, fmt.Errorf("unable to query photos: %v", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Photo])
}

func (r *SharedAlbumRepo) GetPhoto(ctx context.Context, albumId, photoId string) (*model.Photo, error) {
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
	where ap.album_id = $albumId
	and ap.photo_id = $photoId`

	var err error
	var photo model.Photo

	args := pgx.NamedArgs{
		"albumId": albumId,
		"photoId": photoId,
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

package repo

import (
	"context"
	"fmt"
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
	var err error

	query := `select
		id,
		user_id,
		name,
		description,
		cover_photo_id,
		is_share_link_enabled,
		read_access_hash,
		created_at,
		updated_at
	from albums where id = $1`

	var album model.Album

	err = r.db.QueryRow(ctx, query, albumId).Scan(
		&album.Id,
		&album.UserId,
		&album.Name,
		&album.Description,
		&album.CoverPhotoId,
		&album.IsShareLinkEnabled,
		&album.ReadAccessHash,
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

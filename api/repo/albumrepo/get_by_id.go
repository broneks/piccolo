package albumrepo

import (
	"context"
	"fmt"
	"piccolo/api/model"

	"github.com/jackc/pgx/v5"
)

// Checks for read access
func (repo *AlbumRepo) GetById(ctx context.Context, albumId, userId string) (*model.Album, error) {
	var err error

	canRead, err := repo.CanReadAlbum(ctx, albumId, userId)
	if err != nil {
		return nil, err
	}
	if !canRead {
		return nil, fmt.Errorf("unauthorized")
	}

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

	err = repo.db.QueryRow(ctx, query, albumId).Scan(
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

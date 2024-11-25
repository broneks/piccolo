package sharedalbumrepo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Permits:
// - non-users with a valid album read access hash for an album that has sharing enabled
func (repo *SharedAlbumRepo) CanReadSharedAlbum(ctx context.Context, albumId, readAccessHash string) (bool, error) {
	query := `select exists (
		select 1 from albums
		where id = @albumId and is_share_link_enabled = true and read_access_hash = @readAccessHash
	) as can`

	var can bool

	args := pgx.NamedArgs{
		"albumId":        albumId,
		"readAccessHash": readAccessHash,
	}

	err := repo.db.QueryRow(ctx, query, args).Scan(&can)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, fmt.Errorf("no rows found for album id '%s'", albumId)
		}
		return false, fmt.Errorf("query error: %v", err)
	}

	return can, nil
}

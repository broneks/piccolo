package albumrepo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Permits:
// - album owner
// - album user with any role
func (repo *AlbumRepo) CanReadAlbum(ctx context.Context, albumId, userId string) (bool, error) {
	query := `select exists (
		select 1 from albums where id = @albumId and user_id = @userId
		union
		select 1 from album_users where album_id = @albumId and user_id = @userId
	) as can`

	var can bool

	args := pgx.NamedArgs{
		"albumId": albumId,
		"userId":  userId,
	}
	err := repo.db.QueryRow(ctx, query, args).Scan(&can)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, fmt.Errorf("no rows found for album id '%s' and user id '%s'", albumId, userId)
		}
		return false, fmt.Errorf("query error: %v", err)
	}

	return can, nil
}

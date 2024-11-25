package albumrepo

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

// Checks for read access
func (repo *AlbumRepo) GetUserRole(ctx context.Context, albumId, userId string) (string, error) {
	var err error

	canRead, err := repo.CanReadAlbum(ctx, albumId, userId)
	if err != nil {
		return "", err
	}
	if !canRead {
		return "", fmt.Errorf("unauthorized")
	}

	query := `select role from album_users where album_id = @albumId and user_id = @userId`

	var role string

	args := pgx.NamedArgs{
		"albumId": albumId,
		"userId":  userId,
	}
	err = repo.db.QueryRow(ctx, query, args).Scan(&role)
	if err != nil {
		slog.Debug(err.Error())
		if err == pgx.ErrNoRows {
			return "", fmt.Errorf("no rows found for album id '%s' and user id '%s'", albumId, userId)
		}
		return "", fmt.Errorf("query error: %v", err)
	}

	return role, nil
}

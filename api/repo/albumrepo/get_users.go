package albumrepo

import (
	"context"
	"fmt"
	"log/slog"
	"piccolo/api/model"
)

// Checks for read access
func (repo *AlbumRepo) GetUsers(ctx context.Context, albumId, userId string) ([]model.AlbumUserWithUser, error) {
	var err error

	canRead, err := repo.CanReadAlbum(ctx, albumId, userId)
	if err != nil {
		return nil, err
	}
	if !canRead {
		return nil, fmt.Errorf("unauthorized")
	}

	query := `select
		au.album_id,
		au.user_id,
		au.role,
		au.created_at,
		u.id,
		u.username,
		u.email,
		u.hash,
		u.hashed_at,
		u.last_login_at,
		u.created_at,
		u.updated_at
	from users u
	join album_users au on u.id = au.user_id
	where au.album_id = $1
	order by u.created_at desc`

	var albumUsers []model.AlbumUserWithUser

	rows, err := repo.db.Query(ctx, query, albumId)
	if err != nil {
		slog.Debug(err.Error())
		return nil, fmt.Errorf("unable to query users or album users: %v", err)
	}

	for rows.Next() {
		albumUser := new(model.AlbumUserWithUser)

		err = rows.Scan(
			&albumUser.AlbumUser.AlbumId,
			&albumUser.AlbumUser.UserId,
			&albumUser.AlbumUser.Role,
			&albumUser.AlbumUser.CreatedAt,
			&albumUser.User.Id,
			&albumUser.User.Username,
			&albumUser.User.Email,
			&albumUser.User.Hash,
			&albumUser.User.HashedAt,
			&albumUser.User.LastLoginAt,
			&albumUser.User.CreatedAt,
			&albumUser.User.UpdatedAt,
		)
		if err != nil {
			slog.Debug(err.Error())
			return nil, fmt.Errorf("row scan failed: %w", err)
		}

		albumUsers = append(albumUsers, *albumUser)
	}
	if err = rows.Err(); err != nil {
		slog.Debug(err.Error())
		return nil, fmt.Errorf("iteration error: %w", err)
	}

	return albumUsers, nil
}

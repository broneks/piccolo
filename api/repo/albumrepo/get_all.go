package albumrepo

import (
	"context"
	"fmt"
	"log/slog"
	"piccolo/api/model"

	"github.com/jackc/pgx/v5"
)

// Get all scoped to user - albums irrespective of if the user is the owner or just a member
func (repo *AlbumRepo) GetAll(ctx context.Context, userId string) ([]model.Album, error) {
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
	from albums a1 where user_id = $1
	union
	select
		a.id,
		a.user_id,
		a.name,
		a.description,
		a.cover_photo_id,
		a.is_share_link_enabled,
		a.read_access_hash,
		a.created_at,
		a.updated_at
	from albums a join album_users au on a.id = au.album_id where au.user_id = $1
	order by created_at desc`

	rows, err := repo.db.Query(ctx, query, userId)
	if err != nil {
		slog.Debug(err.Error())
		return nil, fmt.Errorf("unable to query albums: %v", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Album])
}

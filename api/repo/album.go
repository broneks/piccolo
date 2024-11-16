package repo

import (
	"context"
	"fmt"
	"log/slog"
	"piccolo/api/model"
	"piccolo/api/types"
	"piccolo/api/util"

	"github.com/jackc/pgx/v5"
)

type AlbumRepo struct {
	db types.ServerDB
}

func NewAlbumRepo(db types.ServerDB) *AlbumRepo {
	return &AlbumRepo{db: db}
}

// Permits:
// - album owner
// - album user with any role
func (r *AlbumRepo) CanReadAlbum(ctx context.Context, albumId, userId string) (bool, error) {
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
	err := r.db.QueryRow(ctx, query, args).Scan(&can)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, fmt.Errorf("no rows found for album id '%s' and user id '%s'", albumId, userId)
		}
		return false, fmt.Errorf("query error: %v", err)
	}

	return can, nil
}

// Permits:
// - album owner
// - album user with editor role
func (r *AlbumRepo) CanWriteAlbum(ctx context.Context, albumId, userId string) (bool, error) {
	query := `select exists (
		select 1 from albums where id = @albumId and user_id = @userId
		union
		select 1 from album_users where album_id = @albumId and user_id = @userId and role in ('editor')
	) as can`

	var can bool

	args := pgx.NamedArgs{
		"albumId": albumId,
		"userId":  userId,
	}
	err := r.db.QueryRow(ctx, query, args).Scan(&can)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, fmt.Errorf("no rows found for album id '%s' and user id '%s'", albumId, userId)
		}
		return false, fmt.Errorf("query error: %v", err)
	}

	return can, nil
}

// Checks for read access
func (r *AlbumRepo) GetById(ctx context.Context, albumId, userId string) (*model.Album, error) {
	var err error

	canRead, err := r.CanReadAlbum(ctx, albumId, userId)
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

// Get all albums irrespective of if the user is the owner or just a member
func (r *AlbumRepo) GetAll(ctx context.Context, userId string) ([]model.Album, error) {
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
	from albums a join album_users au on a.id = au.album_id where au.user_id = $1`

	rows, err := r.db.Query(ctx, query, userId)
	if err != nil {
		slog.Debug(err.Error())
		return nil, fmt.Errorf("unable to query albums: %v", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Album])
}

func (r *AlbumRepo) GetUserRole(ctx context.Context, albumId, userId string) (string, error) {
	query := `select role from album_users where album_id = $albumId and user_id = $userId`

	var role string

	args := pgx.NamedArgs{
		"albumId": albumId,
		"userId":  userId,
	}
	err := r.db.QueryRow(ctx, query, args).Scan(&role)
	if err != nil {
		slog.Debug(err.Error())
		if err == pgx.ErrNoRows {
			return "", fmt.Errorf("no rows found for album id '%s' and user id '%s'", albumId, userId)
		}
		return "", fmt.Errorf("query error: %v", err)
	}

	return role, nil
}

// Checks for read access
func (r *AlbumRepo) GetUsers(ctx context.Context, albumId, userId string) ([]model.AlbumUserWithUser, error) {
	var err error

	canRead, err := r.CanReadAlbum(ctx, albumId, userId)
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
	where au.album_id = $1`

	var albumUsers []model.AlbumUserWithUser

	rows, err := r.db.Query(ctx, query, albumId)
	if err != nil {
		slog.Debug(err.Error())
		return nil, fmt.Errorf("unable to query users or album users: %v", err)
	}

	for rows.Next() {
		albumUser := new(model.AlbumUserWithUser)

		err = rows.Scan(
			&albumUser.AlbumId,
			&albumUser.UserId,
			&albumUser.Role,
			&albumUser.CreatedAt,
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

// Checks for read access
func (r *AlbumRepo) GetPhotos(ctx context.Context, albumId, userId string) ([]model.Photo, error) {
	var err error

	canRead, err := r.CanReadAlbum(ctx, albumId, userId)
	if err != nil {
		return nil, err
	}
	if !canRead {
		return nil, fmt.Errorf("unauthorized")
	}

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
	where ap.album_id = $1`

	rows, err := r.db.Query(ctx, query, albumId)
	if err != nil {
		slog.Debug(err.Error())
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
		cover_photo_id,
		is_share_link_enabled
	) values (
		@userId,
		@name,
		@description,
		@coverPhotoId,
		@isShareLinkEnabled
	)`

	args := pgx.NamedArgs{
		"userId":             album.UserId,
		"name":               album.Name,
		"description":        album.Description,
		"coverPhotoId":       album.CoverPhotoId,
		"isShareLinkEnabled": album.IsShareLinkEnabled,
	}
	_, err := r.db.Exec(ctx, query, args)
	if err != nil {
		slog.Debug(err.Error())
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

// TODO: control which fields can be updated
// Checks for write access
func (r *AlbumRepo) Update(ctx context.Context, album model.Album, userId string) error {
	return nil
}

// Checks for write access
func (r *AlbumRepo) InsertUsers(ctx context.Context, albumId string, albumUsers []model.AlbumUser, userId string) error {
	var err error

	canWrite, err := r.CanWriteAlbum(ctx, albumId, userId)
	if err != nil {
		return err
	}
	if !canWrite {
		return fmt.Errorf("unauthorized")
	}

	query := `insert into album_users (
		album_id,
		user_id,
		role
	) values (
		@albumId,
		@userId,
		@role
	)`

	batch := &pgx.Batch{}

	for _, albumUser := range albumUsers {
		args := pgx.NamedArgs{
			"albumId": albumId,
			"userId":  albumUser.UserId,
			"role":    albumUser.Role,
		}
		batch.Queue(query, args)
	}

	results := r.db.SendBatch(ctx, batch)
	defer results.Close()

	for _, albumUser := range albumUsers {
		_, err := results.Exec()
		if err != nil {
			if sqlErr := util.CheckSqlError(err); sqlErr == "unique-violation" {
				slog.Debug(fmt.Sprintf("album user \"%s\" already exists", albumUser.UserId.String))
				continue
			}

			slog.Debug(err.Error())
			return fmt.Errorf("unable to insert album user: %w", err)
		}
	}

	return nil
}

// TODO
// Checks for write access
func (r *AlbumRepo) UpdateUsers(ctx context.Context, usersToUpdate []model.AlbumUser, userId string) error {
	return nil
}

// TODO
// Checks for write access
func (r *AlbumRepo) RemoveUsers(ctx context.Context, userIdsToRemove []string, userId string) error {
	return nil
}

// Checks for write access
func (r *AlbumRepo) InsertPhotos(ctx context.Context, albumId string, photoIds []string, userId string) error {
	var err error

	canWrite, err := r.CanWriteAlbum(ctx, albumId, userId)
	if err != nil {
		return err
	}
	if !canWrite {
		return fmt.Errorf("unauthorized")
	}

	query := `insert into album_photos (
		album_id,
		photo_id,
		user_id
	) values (
		@albumId,
		@photoId,
		@userId
	)`

	batch := &pgx.Batch{}

	for _, photoId := range photoIds {
		args := pgx.NamedArgs{
			"albumId": albumId,
			"photoId": photoId,
			"userId":  userId,
		}
		batch.Queue(query, args)
	}

	results := r.db.SendBatch(ctx, batch)
	defer results.Close()

	for _, photoId := range photoIds {
		_, err := results.Exec()
		if err != nil {
			if sqlErr := util.CheckSqlError(err); sqlErr == "unique-violation" {
				slog.Debug(fmt.Sprintf("photo \"%s\" already exists", photoId))
				continue
			}

			slog.Debug(err.Error())
			return fmt.Errorf("unable to insert album photo: %w", err)
		}
	}

	return nil
}

// TODO
// Checks for write access
func (r *AlbumRepo) RemovePhotos(ctx context.Context, photoIds []string, userId string) error {
	return nil
}

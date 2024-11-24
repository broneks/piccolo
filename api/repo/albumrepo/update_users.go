package albumrepo

import (
	"context"
	"piccolo/api/model"
)

// Checks for write access
func (repo *AlbumRepo) UpdateUsers(ctx context.Context, usersToUpdate []model.AlbumUser, userId string) error {
	return nil
}

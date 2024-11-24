package albumrepo

import (
	"context"
	"piccolo/api/model"
)

// TODO: control which fields can be updated
// Checks for write access
func (repo *AlbumRepo) Update(ctx context.Context, album model.Album, userId string) error {
	return nil
}

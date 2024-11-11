package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type AlbumPhoto struct {
	AlbumId pgtype.UUID `json:"albumId"`
	PhotoId pgtype.UUID `json:"photoId"`
}

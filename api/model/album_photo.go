package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type AlbumPhoto struct {
	AlbumId   pgtype.Text        `json:"albumId"`
	PhotoId   pgtype.Text        `json:"photoId"`
	UserId    pgtype.Text        `json:"userId"`
	CreatedAt pgtype.Timestamptz `json:"createdAt"`
}

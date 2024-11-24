package model

import "github.com/jackc/pgx/v5/pgtype"

type AlbumPhotoFavourites struct {
	PhotoId   pgtype.Text        `json:"photoId"`
	CreatedAt pgtype.Timestamptz `json:"createdAt"`
}

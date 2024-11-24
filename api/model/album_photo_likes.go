package model

import "github.com/jackc/pgx/v5/pgtype"

type AlbumPhotoLikes struct {
	PhotoId pgtype.Text `json:"photoId"`
	Likes   pgtype.Int4 `json:"likes"` // aggregate total
}

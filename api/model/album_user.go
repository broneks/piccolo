package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type AlbumUser struct {
	AlbumId   pgtype.Text        `json:"albumId"`
	UserId    pgtype.Text        `json:"userId"`
	Role      pgtype.Text        `json:"role"`
	CreatedAt pgtype.Timestamptz `json:"createdAt"`
}

type AlbumUserWithUser struct {
	AlbumId   pgtype.Text        `json:"albumId"`
	UserId    pgtype.Text        `json:"userId"`
	Role      pgtype.Text        `json:"role"`
	CreatedAt pgtype.Timestamptz `json:"createdAt"`
	User      `json:"user"`
}

func NewAlbumUserWithUser() *AlbumUserWithUser {
	return &AlbumUserWithUser{
		User: User{},
	}
}

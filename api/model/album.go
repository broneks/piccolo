package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Album struct {
	Id             pgtype.Text        `json:"id"`
	UserId         pgtype.Text        `json:"userId"`
	Name           pgtype.Text        `json:"name"`
	Description    pgtype.Text        `json:"description"`
	CoverPhotoId   pgtype.Text        `json:"coverPhotoId"`
	ReadAccessHash pgtype.Text        `json:"readAccessHash"`
	CreatedAt      pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt      pgtype.Timestamptz `json:"-"`
}

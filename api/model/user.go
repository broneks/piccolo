package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	Id          pgtype.Text        `json:"id"`
	Username    pgtype.Text        `json:"username"`
	Email       pgtype.Text        `json:"email"`
	Hash        pgtype.Text        `json:"-"`
	HashedAt    pgtype.Timestamptz `json:"-"`
	LastLoginAt pgtype.Timestamptz `json:"-"`
	CreatedAt   pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt   pgtype.Timestamptz `json:"-"`
}

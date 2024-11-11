package model

import (
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
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

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Hash.String), []byte(password))
	return err == nil
}

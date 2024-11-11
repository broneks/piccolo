package model

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id          string         `json:"id"`
	Username    sql.NullString `json:"username"`
	Email       string         `json:"email"`
	Hash        string         `json:"-"`
	HashedAt    time.Time      `json:"-"`
	LastLoginAt time.Time      `json:"-"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"-"`
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Hash), []byte(password))
	return err == nil
}

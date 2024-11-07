package model

import "time"

type User struct {
	Id          string    `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Hash        string    `json:"-"`
	HashedAt    time.Time `json:"-"`
	LastLoginAt time.Time `json:"-"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"-"`
}

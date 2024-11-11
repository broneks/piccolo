package model

import (
	"time"
)

type Album struct {
	Id           string    `json:"id"`
	UserId       string    `json:"userId"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CoverPhotoId int       `json:"coverPhotoId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"-"`
}

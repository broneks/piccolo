package model

import (
	"database/sql"
	"time"
)

type Album struct {
	Id           string         `json:"id"`
	UserId       string         `json:"userId"`
	Name         string         `json:"name"`
	Description  sql.NullString `json:"description"`
	CoverPhotoId sql.NullString `json:"coverPhotoId"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"-"`
}

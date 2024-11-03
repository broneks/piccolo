package model

import "time"

type Photo struct {
	Id        string
	Location  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

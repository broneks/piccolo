package model

import "time"

type Photo struct {
	Id          string
	Location    string
	Filename    string
	FileSize    int
	ContentType string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

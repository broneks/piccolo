package model

import (
	"context"
	"piccolo/api/storage/wasabi"
	"time"
)

type Photo struct {
	Id          string
	Location    string
	Filename    string
	FileSize    int
	ContentType string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (p *Photo) GetUrl() string {
	url, _ := wasabi.GetPresignedUrl(context.Background(), p.Filename)

	return url
}

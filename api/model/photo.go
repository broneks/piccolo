package model

import (
	"context"
	"piccolo/api/shared"
	"time"
)

type Photo struct {
	Id          string    `json:"id"`
	UserId      string    `json:"userId"`
	Location    string    `json:"-"`
	Filename    string    `json:"filename"`
	FileSize    int       `json:"fileSize"`
	ContentType string    `json:"contentType"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"-"`
}

func (p *Photo) GetUrl(ctx context.Context, server *shared.Server) string {
	key := p.Id

	val, err := server.Cache.Get(ctx, key)
	if err != nil {
		server.Logger.Error(err.Error())
		return ""
	}

	if val != "" {
		return val
	} else {
		url, expirationDuration := server.ObjectStorage.GetPresignedUrl(context.Background(), p.Filename)

		err := server.Cache.Set(ctx, key, url, expirationDuration-(time.Minute*5))
		if err != nil {
			server.Logger.Error(err.Error())
		}

		return url
	}
}

package model

import (
	"context"
	"piccolo/api/shared"
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

func (p *Photo) GetUrl(ctx context.Context, server *shared.Server) string {
	key := p.Id

	val, err := server.Redis.Get(ctx, key)
	if err != nil {
		server.Logger.Error(err.Error())
		return ""
	}

	if val != "" {
		return val
	} else {
		url, expirationDuration := server.Wasabi.GetPresignedUrl(context.Background(), p.Filename)

		err := server.Redis.Set(ctx, key, url, expirationDuration-(time.Minute*5))
		if err != nil {
			server.Logger.Error(err.Error())
		}

		return url
	}
}

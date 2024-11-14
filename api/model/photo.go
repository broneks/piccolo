package model

import (
	"context"
	"piccolo/api/types"

	"github.com/jackc/pgx/v5/pgtype"
)

type Photo struct {
	Id          pgtype.Text        `json:"id"`
	UserId      pgtype.Text        `json:"userId"`
	Location    pgtype.Text        `json:"-"`
	Filename    pgtype.Text        `json:"filename"`
	FileSize    pgtype.Int4        `json:"fileSize"`
	ContentType pgtype.Text        `json:"contentType"`
	CreatedAt   pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt   pgtype.Timestamptz `json:"-"`
}

func (p *Photo) GetUrl(ctx context.Context, server *types.Server) string {
	key := p.Id.String

	val, err := server.Cache.Get(ctx, key)
	if err != nil {
		server.Logger.Error(err.Error())
		return ""
	}

	if val != "" {
		return val
	} else {
		url, expirationDuration := server.ObjectStorage.GetPresignedUrl(context.Background(), p.Filename.String, p.UserId.String)

		err := server.Cache.Set(ctx, key, url, expirationDuration)
		if err != nil {
			server.Logger.Error(err.Error())
		}

		return url
	}
}

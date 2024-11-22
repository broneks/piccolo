package model

import (
	"context"
	"fmt"
	"piccolo/api/consts"
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

type PhotoWithUrl struct {
	*Photo
	Url string `json:"url"`
}

func (photo *Photo) GetUrl(ctx context.Context, server *types.Server) string {
	key := fmt.Sprintf(consts.PhotoPresignedUrlKeyFormat, photo.Id.String)

	val, err := server.Cache.Get(ctx, key)
	if err != nil {
		server.Logger.Error(err.Error())
		return ""
	}

	if val != "" {
		return val
	} else {
		url, expirationDuration := server.ObjectStorage.GetPresignedUrl(context.Background(), photo.Filename.String, photo.UserId.String)

		err := server.Cache.Set(ctx, key, url, expirationDuration)
		if err != nil {
			server.Logger.Error(err.Error())
		}

		return url
	}
}

func NewPhotoWithUrl(ctx context.Context, server *types.Server, photo *Photo) *PhotoWithUrl {
	return &PhotoWithUrl{
		Photo: photo,
		Url:   photo.GetUrl(ctx, server),
	}
}

func NewPhotosWithUrl(ctx context.Context, server *types.Server, photos []Photo) []*PhotoWithUrl {
	var photosWithUrl []*PhotoWithUrl

	for _, photo := range photos {
		photosWithUrl = append(photosWithUrl, &PhotoWithUrl{
			Photo: &photo,
			Url:   photo.GetUrl(ctx, server),
		})
	}

	return photosWithUrl
}

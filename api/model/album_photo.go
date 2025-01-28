package model

import (
	"context"
	"piccolo/api/types"

	"github.com/jackc/pgx/v5/pgtype"
)

type AlbumPhoto struct {
	Photo
	AddedAt pgtype.Timestamptz `json:"addedAt"`
}

type AlbumPhotoWithUrl struct {
	*AlbumPhoto
	Url string `json:"url"`
}

func NewAlbumPhotoWithUrl(ctx context.Context, server *types.Server, albumPhoto *AlbumPhoto) *AlbumPhotoWithUrl {
	return &AlbumPhotoWithUrl{
		AlbumPhoto: albumPhoto,
		Url:        albumPhoto.Photo.GetUrl(ctx, server),
	}
}

func NewAlbumPhotosWithUrl(ctx context.Context, server *types.Server, albumPhotos []AlbumPhoto) []*AlbumPhotoWithUrl {
	var albumPhotosWithUrl []*AlbumPhotoWithUrl

	for _, albumPhoto := range albumPhotos {
		albumPhotosWithUrl = append(albumPhotosWithUrl, NewAlbumPhotoWithUrl(ctx, server, &albumPhoto))
	}

	return albumPhotosWithUrl
}

package page

import (
	"log/slog"
	"net/http"
	"piccolo/api/helper"
	"piccolo/api/model"
	"piccolo/api/repo/sharedalbumrepo"
	"piccolo/api/types"
	"sync"

	"github.com/labstack/echo/v4"
)

type SharedAlbumPayload struct {
	PageInfo
	*model.Album
	CoverPhoto *model.AlbumPhotoWithUrl
	Photos     []*model.AlbumPhotoWithUrl
}

func handleSharedAlbumPage(server *types.Server, sharedAlbumRepo *sharedalbumrepo.SharedAlbumRepo) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		albumId := helper.GetIdParam(c)

		var wg sync.WaitGroup
		wg.Add(2)

		var album *model.Album
		var photos []model.AlbumPhoto
		var albumErr, photosErr error

		go func() {
			defer wg.Done()
			album, albumErr = sharedAlbumRepo.GetById(ctx, albumId)
		}()

		go func() {
			defer wg.Done()
			photos, photosErr = sharedAlbumRepo.GetPhotos(ctx, albumId)
		}()

		wg.Wait()

		if albumErr != nil || photosErr != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		photosWithUrl := model.NewAlbumPhotosWithUrl(ctx, server, photos)

		var coverPhoto *model.AlbumPhotoWithUrl

		if album.CoverPhotoId.String != "" {
			photo, err := sharedAlbumRepo.GetPhoto(ctx, albumId, album.CoverPhotoId.String)
			if err != nil {
				slog.Debug("failed to get photo", "err", err)
			} else {
				coverPhoto = model.NewAlbumPhotoWithUrl(ctx, server, photo)
			}
		}

		pageInfo := NewPageInfo(c, album.Name.String)

		return c.Render(http.StatusOK, "shared_album.html", &SharedAlbumPayload{
			PageInfo:   pageInfo,
			Album:      album,
			CoverPhoto: coverPhoto,
			Photos:     photosWithUrl,
		})
	}
}

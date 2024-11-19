package page

import (
	"net/http"
	"piccolo/api/model"
	"piccolo/api/repo"
	"piccolo/api/types"
	"piccolo/api/util"
	"sync"

	"github.com/labstack/echo/v4"
)

type Payload struct {
	*model.Album
	CoverPhoto *model.PhotoWithUrl
	Photos     []*model.PhotoWithUrl
}

func handleSharedAlbumPage(server *types.Server, sharedAlbumRepo *repo.SharedAlbumRepo) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		albumId := util.GetIdParam(c)

		var wg sync.WaitGroup
		wg.Add(2)

		var album *model.Album
		var photos []model.Photo
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

		photosWithUrl := model.NewPhotosWithUrl(ctx, server, photos)

		var coverPhoto *model.PhotoWithUrl

		if album.CoverPhotoId.String != "" {
			photo, err := sharedAlbumRepo.GetPhoto(ctx, albumId, album.CoverPhotoId.String)
			if err != nil {
				server.Logger.Debug(err.Error())
			} else {
				coverPhoto = model.NewPhotoWithUrl(ctx, server, photo)
			}
		}

		return c.Render(http.StatusOK, "shared_album.html", &Payload{
			Album:      album,
			CoverPhoto: coverPhoto,
			Photos:     photosWithUrl,
		})
	}
}

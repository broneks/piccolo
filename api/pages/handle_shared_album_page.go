package pages

import (
	"net/http"
	"piccolo/api/model"
	"piccolo/api/repo"
	"piccolo/api/types"
	"piccolo/api/util"

	"github.com/labstack/echo/v4"
)

type Payload struct {
	*model.Album
	Photos []*model.PhotoWithUrl
}

func handleSharedAlbumPage(server *types.Server, sharedAlbumRepo *repo.SharedAlbumRepo) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		albumId := util.GetIdParam(c)
		album, err := sharedAlbumRepo.GetById(ctx, albumId)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		photos, err := sharedAlbumRepo.GetPhotos(ctx, albumId)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		photosWithUrl := model.NewPhotosWithUrl(ctx, server, photos)

		return c.Render(http.StatusOK, "album.html", &Payload{
			Album:  album,
			Photos: photosWithUrl,
		})
	}
}

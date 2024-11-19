package album

import (
	"net/http"
	"piccolo/api/model"
	"piccolo/api/types"
	"piccolo/api/util"

	"github.com/labstack/echo/v4"
)

func (m *AlbumModule) getAlbumPhotosHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	albumId := util.GetIdParam(c)
	if albumId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id param.")
	}

	photos, err := m.albumRepo.GetPhotos(ctx, albumId, userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, types.SuccessRes{
			Success: false,
			Message: "Not found",
		})
	}

	if len(photos) == 0 {
		return c.JSON(http.StatusOK, []any{})
	}

	photosWithUrl := model.NewPhotosWithUrl(ctx, m.server, photos)

	return c.JSON(http.StatusOK, photosWithUrl)
}

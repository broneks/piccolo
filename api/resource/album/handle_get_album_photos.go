package album

import (
	"net/http"
	"piccolo/api/helper"
	"piccolo/api/model"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func (mod *AlbumModule) getAlbumPhotosHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	albumId := helper.GetIdParam(c)
	if albumId == "" {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Invalid id param.",
		})
	}

	photos, err := mod.albumRepo.GetPhotos(ctx, albumId, userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, types.SuccessRes{
			Success: false,
			Message: "Not found",
		})
	}

	if len(photos) == 0 {
		return c.JSON(http.StatusOK, []any{})
	}

	photosWithUrl := model.NewPhotosWithUrl(ctx, mod.server, photos)

	return c.JSON(http.StatusOK, photosWithUrl)
}

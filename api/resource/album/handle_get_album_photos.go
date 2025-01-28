package album

import (
	"log/slog"
	"net/http"
	"piccolo/api/helper"
	"piccolo/api/model"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func (mod *AlbumModule) getAlbumPhotosHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)
	queryParams := new(types.ListQueryParams)

	var err error

	albumId := helper.GetIdParam(c)
	if albumId == "" {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Invalid id param.",
		})
	}

	if err = c.Bind(queryParams); err != nil {
		slog.Error("failed to bind get album photos query params", "err", err)

		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: err.Error(),
		})
	}

	if err = c.Validate(queryParams); err != nil {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: err.Error(),
		})
	}

	photos, err := mod.albumRepo.GetPhotosWithParams(ctx, albumId, userId, *queryParams)
	if err != nil {
		slog.Debug("failed to get album photos", "err", err)
		return c.JSON(http.StatusNotFound, types.SuccessRes{
			Success: false,
			Message: "Not found",
		})
	}

	if len(photos) == 0 {
		return c.JSON(http.StatusOK, []any{})
	}

	photosWithUrl := model.NewAlbumPhotosWithUrl(ctx, mod.server, photos)

	return c.JSON(http.StatusOK, photosWithUrl)
}

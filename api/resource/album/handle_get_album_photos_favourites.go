package album

import (
	"net/http"
	"piccolo/api/helper"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func (mod *AlbumModule) getAlbumPhotosFavouritesHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	albumId := helper.GetIdParam(c)
	if albumId == "" {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Invalid id param.",
		})
	}

	photosFavourites, err := mod.albumRepo.GetPhotosFavourites(ctx, albumId, userId)
	if err != nil {
		if err.Error() == "unauthorized" {
			return c.JSON(http.StatusNotFound, types.SuccessRes{
				Success: false,
				Message: "Not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, types.SuccessRes{
			Success: false,
			Message: "Unexpected error",
		})
	}

	if len(photosFavourites) == 0 {
		return c.JSON(http.StatusOK, []any{})
	}

	return c.JSON(http.StatusOK, photosFavourites)
}

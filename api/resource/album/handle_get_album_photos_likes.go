package album

import (
	"net/http"
	"piccolo/api/helper"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func (mod *AlbumModule) getAlbumPhotosLikesHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	albumId := helper.GetIdParam(c)
	if albumId == "" {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Invalid id param.",
		})
	}

	// TODO cache likes
	photoIds := make([]string, 0)
	// photoIds := helper.GetListParam(c, "photoIds")

	photosLikes, err := mod.albumRepo.GetPhotosLikes(ctx, albumId, photoIds, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.SuccessRes{
			Success: false,
			Message: "Unexpected error",
		})
	}

	if len(photosLikes) == 0 {
		return c.JSON(http.StatusOK, []any{})
	}

	return c.JSON(http.StatusOK, photosLikes)
}

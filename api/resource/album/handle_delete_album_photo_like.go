package album

import (
	"net/http"
	"piccolo/api/helper"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func (mod *AlbumModule) deleteAlbumPhotoLike(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	albumId := helper.GetIdParam(c)
	if albumId == "" {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Invalid album id param.",
		})
	}

	photoId := helper.GetIdParamByName(c, "photoId")
	if photoId == "" {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Invalid photo id param.",
		})
	}

	rowsAffected, err := mod.albumRepo.UnlikePhoto(ctx, albumId, photoId, userId)
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

	if rowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Photo is not liked",
		})
	}

	return c.JSON(http.StatusOK, types.SuccessRes{
		Success: true,
		Message: "Unliked photo",
	})
}

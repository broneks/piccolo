package album

import (
	"log/slog"
	"net/http"
	"piccolo/api/helper"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func (mod *AlbumModule) deleteAlbumPhotoHandler(c echo.Context) error {
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

	rowsAffected, err := mod.albumRepo.RemovePhotoOne(ctx, albumId, photoId, userId)
	if err != nil {
		if err.Error() == "unauthorized" {
			return c.JSON(http.StatusNotFound, types.SuccessRes{
				Success: false,
				Message: "Not found",
			})
		}

		slog.Error("error removing album photo", "err", err)
		return c.JSON(http.StatusInternalServerError, types.SuccessRes{
			Success: false,
			Message: "Unexpected error",
		})
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Album photo is not found",
		})
	}

	return c.JSON(http.StatusOK, types.SuccessRes{
		Success: true,
		Message: "Deleted album photo",
	})
}

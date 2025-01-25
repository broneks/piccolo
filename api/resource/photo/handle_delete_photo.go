package photo

import (
	"net/http"
	"piccolo/api/helper"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func (mod *PhotoModule) deletePhotoHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	photoId := helper.GetIdParam(c)
	if photoId == "" {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Invalid id param.",
		})
	}

	rowsAffected, err := mod.photoRepo.RemoveOne(ctx, photoId, userId)
	if err != nil || rowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Photo not found",
		})
	}

	return c.JSON(http.StatusOK, types.SuccessRes{
		Success: true,
		Message: "Deleted photo",
	})
}

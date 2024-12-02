package photo

import (
	"net/http"
	"piccolo/api/helper"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func (mod *PhotoModule) deletePhotosHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	photoIds := helper.GetListParam(c, "ids")
	if len(photoIds) == 0 {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Missing ids param.",
		})
	}

	rowsAffected, err := mod.photoRepo.RemoveMany(ctx, photoIds, userId)
	if err != nil || rowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Photo not found",
		})
	}

	return c.JSON(http.StatusOK, types.SuccessRes{
		Success: true,
		Message: "Deleted photos",
	})
}

package photo

import (
	"net/http"
	"piccolo/api/helper"
	"piccolo/api/model"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func (mod *PhotoModule) getPhotoHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	photoId := helper.GetIdParam(c)
	if photoId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id param.")
	}

	photo, err := mod.photoRepo.GetById(ctx, photoId, userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, types.SuccessRes{
			Success: false,
			Message: "Not found",
		})
	}

	photoWithUrl := model.NewPhotoWithUrl(ctx, mod.server, photo)

	return c.JSON(http.StatusOK, photoWithUrl)
}

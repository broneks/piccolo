package photo

import (
	"net/http"
	"piccolo/api/model"
	"piccolo/api/types"
	"piccolo/api/util"

	"github.com/labstack/echo/v4"
)

func (m *PhotoModule) getPhotoHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	photoId := util.GetIdParam(c)
	if photoId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id param.")
	}

	photo, err := m.photoRepo.GetById(ctx, photoId, userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, types.SuccessRes{
			Success: false,
			Message: "Not found",
		})
	}

	photoWithUrl := model.NewPhotoWithUrl(ctx, m.server, photo)

	return c.JSON(http.StatusOK, photoWithUrl)
}

package photo

import (
	"net/http"
	"piccolo/api/model"

	"github.com/labstack/echo/v4"
)

func (m *PhotoModule) getPhotosHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	photos, _ := m.photoRepo.GetAll(ctx, userId)

	if len(photos) == 0 {
		return c.JSON(http.StatusOK, []any{})
	}

	photosWithUrl := model.NewPhotosWithUrl(ctx, m.server, photos)

	return c.JSON(http.StatusOK, photosWithUrl)
}

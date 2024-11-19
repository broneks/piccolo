package photo

import (
	"net/http"
	"piccolo/api/model"

	"github.com/labstack/echo/v4"
)

func (mod *PhotoModule) getPhotosHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	photos, _ := mod.photoRepo.GetAll(ctx, userId)

	if len(photos) == 0 {
		return c.JSON(http.StatusOK, []any{})
	}

	photosWithUrl := model.NewPhotosWithUrl(ctx, mod.server, photos)

	return c.JSON(http.StatusOK, photosWithUrl)
}

package photo

import (
	"log/slog"
	"net/http"
	"piccolo/api/model"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func (mod *PhotoModule) getPhotosHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)
	queryParams := new(types.ListQueryParams)

	var err error

	if err = c.Bind(queryParams); err != nil {
		slog.Error("failed to bind get photos query params", "err", err)

		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: err.Error(),
		})
	}

	if err = c.Validate(queryParams); err != nil {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: err.Error(),
		})
	}

	photos, _ := mod.photoRepo.GetAllWithParams(ctx, userId, *queryParams)

	if len(photos) == 0 {
		return c.JSON(http.StatusOK, []any{})
	}

	photosWithUrl := model.NewPhotosWithUrl(ctx, mod.server, photos)

	return c.JSON(http.StatusOK, photosWithUrl)
}

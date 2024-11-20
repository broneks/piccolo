package photo

import (
	"net/http"
	"piccolo/api/helper"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func (mod *PhotoModule) getPhotoAlbumsHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	photoId := helper.GetIdParam(c)
	if photoId == "" {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Invalid id param.",
		})
	}

	albums, err := mod.photoRepo.GetAlbums(ctx, photoId, userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, types.SuccessRes{
			Success: false,
			Message: "Not found",
		})
	}

	if len(albums) == 0 {
		return c.JSON(http.StatusOK, []any{})
	}

	return c.JSON(http.StatusOK, albums)
}

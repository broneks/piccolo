package photos

import (
	"net/http"
	"piccolo/api/shared"
	"piccolo/api/util"

	"github.com/labstack/echo/v4"
)

func (m *PhotosModule) getPhotoAlbumsHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	photoId := util.GetIdParam(c)
	if photoId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id param.")
	}

	albums, err := m.photoRepo.GetAlbums(ctx, photoId, userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, shared.SuccessRes{
			Success: false,
			Message: "Not found",
		})
	}

	if len(albums) == 0 {
		return c.JSON(http.StatusOK, []any{})
	}

	return c.JSON(http.StatusOK, albums)
}

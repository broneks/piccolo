package albums

import (
	"net/http"
	"piccolo/api/shared"
	"piccolo/api/util"

	"github.com/labstack/echo/v4"
)

func (m *AlbumsModule) getAlbumHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	albumId := util.GetIdParam(c)
	if albumId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id param.")
	}

	album, err := m.albumRepo.GetById(ctx, albumId, userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, shared.SuccessRes{
			Success: false,
			Message: "Not found",
		})
	}

	return c.JSON(http.StatusOK, album)
}

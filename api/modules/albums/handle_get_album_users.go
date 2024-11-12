package albums

import (
	"net/http"
	"piccolo/api/shared"
	"piccolo/api/util"

	"github.com/labstack/echo/v4"
)

func (m *AlbumsModule) getAlbumUsersHandler(c echo.Context) error {
	ctx := c.Request().Context()

	albumId := util.GetIdParam(c)
	if albumId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id param.")
	}

	users, _ := m.albumRepo.GetUsers(ctx, albumId)

	if len(users) == 0 {
		return c.JSON(http.StatusOK, shared.EmptySlice{})
	}

	return c.JSON(http.StatusOK, users)
}

package album

import (
	"net/http"
	"piccolo/api/types"
	"piccolo/api/util"

	"github.com/labstack/echo/v4"
)

func (m *AlbumModule) getAlbumUsersHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	albumId := util.GetIdParam(c)
	if albumId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id param.")
	}

	albumUsers, err := m.albumRepo.GetUsers(ctx, albumId, userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, types.SuccessRes{
			Success: false,
			Message: "Not found",
		})
	}

	if len(albumUsers) == 0 {
		return c.JSON(http.StatusOK, []any{})
	}

	return c.JSON(http.StatusOK, albumUsers)
}
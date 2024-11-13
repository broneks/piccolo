package albums

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m *AlbumsModule) getAlbumsHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	albums, _ := m.albumRepo.GetAll(ctx, userId)

	if len(albums) == 0 {
		return c.JSON(http.StatusOK, []any{})
	}

	return c.JSON(http.StatusOK, albums)
}

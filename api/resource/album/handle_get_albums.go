package album

import (
	"log/slog"
	"net/http"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func (mod *AlbumModule) getAlbumsHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)
	queryParams := new(types.ListQueryParams)

	var err error

	if err = c.Bind(queryParams); err != nil {
		slog.Error("failed to bind get albums query params", "err", err)

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

	albums, _ := mod.albumRepo.GetAllWithParams(ctx, userId, *queryParams)

	if len(albums) == 0 {
		return c.JSON(http.StatusOK, []any{})
	}

	return c.JSON(http.StatusOK, albums)
}

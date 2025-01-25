package me

import (
	"log/slog"
	"net/http"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func (mod *MeModule) userFileStorageHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	userFileStorage, err := mod.photoService.GetUserFileStorage(ctx, userId)
	if err != nil {
		slog.Error("error getting user file storage", "err", err)

		return c.JSON(http.StatusInternalServerError, types.SuccessRes{
			Success: false,
			Message: "Unexpected error occurred",
		})
	}

	return c.JSON(http.StatusOK, &userFileStorage)
}

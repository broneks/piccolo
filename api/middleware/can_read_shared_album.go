package middleware

import (
	"log/slog"
	"net/http"
	"piccolo/api/helper"
	"piccolo/api/repo/sharedalbumrepo"

	"github.com/labstack/echo/v4"
)

func CanReadSharedAlbum(sharedAlbumRepo *sharedalbumrepo.SharedAlbumRepo) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			albumId := helper.GetIdParam(c)
			if albumId == "" {
				return echo.NewHTTPError(http.StatusNotFound)
			}

			ctx := c.Request().Context()
			readAccessHash := c.QueryParam("share")

			canRead, err := sharedAlbumRepo.CanReadSharedAlbum(ctx, albumId, readAccessHash)
			if err != nil {
				slog.Debug(err.Error())
			}

			if !canRead {
				return echo.NewHTTPError(http.StatusNotFound)
			}

			return next(c)
		}
	}
}

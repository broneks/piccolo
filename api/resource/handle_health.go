package resource

import (
	"log/slog"
	"net/http"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func handleHealth(server *types.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var err error

		if err = server.DB.Ping(ctx); err != nil {
			slog.Error("postgres ping failed", "err", err)
			return c.NoContent(http.StatusServiceUnavailable)
		}

		if err = server.Cache.Ping(ctx); err != nil {
			slog.Error("redis ping failed", "err", err)
			return c.NoContent(http.StatusServiceUnavailable)
		}

		if err = server.ObjectStorage.Ping(ctx); err != nil {
			slog.Error("backblaze ping failed", "err", err)
			return c.NoContent(http.StatusServiceUnavailable)
		}

		return c.NoContent(http.StatusOK)
	}
}

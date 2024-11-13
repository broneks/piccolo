package modules

import (
	"fmt"
	"net/http"
	"piccolo/api/shared"

	"github.com/labstack/echo/v4"
)

func handleHealth(server *shared.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var err error

		if err = server.DB.Ping(ctx); err != nil {
			server.Logger.Error(fmt.Sprintf("postgres ping failed: %v", err.Error()))
			return c.NoContent(http.StatusServiceUnavailable)
		}

		if err = server.Cache.Ping(ctx); err != nil {
			server.Logger.Error(fmt.Sprintf("redis ping failed: %v", err.Error()))
			return c.NoContent(http.StatusServiceUnavailable)
		}

		if err = server.ObjectStorage.Ping(ctx); err != nil {
			server.Logger.Error(fmt.Sprintf("wasabi ping failed: %v", err.Error()))
			return c.NoContent(http.StatusServiceUnavailable)
		}

		return c.NoContent(http.StatusOK)
	}
}

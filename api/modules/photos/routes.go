package photos

import (
	"piccolo/api/middleware"

	"github.com/labstack/echo/v4"
)

func (m *PhotosModule) Routes(g *echo.Group) {
	photos := g.Group("/photos", middleware.Auth())

	photos.GET("", m.getPhotosHandler)
}
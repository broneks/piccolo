package upload

import (
	"piccolo/api/middleware"

	"github.com/labstack/echo/v4"
)

// TODO: move under photos module?
func (m *UploadModule) Routes(g *echo.Group) {
	upload := g.Group("/upload", middleware.Auth())

	upload.POST("", m.postUploadHandler)
}

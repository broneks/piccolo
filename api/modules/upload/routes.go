package upload

import (
	"github.com/labstack/echo/v4"
)

func (m *UploadModule) Routes(g *echo.Group) {
	g.GET("/uploads", m.getUploadsHandler)
	g.POST("/upload", m.postUploadHandler)
}

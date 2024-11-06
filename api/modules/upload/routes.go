package upload

import (
	"github.com/labstack/echo/v4"
)

func (m *UploadModule) Routes(g *echo.Group) {
	g.GET("/uploads", m.getUploadsHandler)
	// e.GET("/uploads", m.handleGetUploads)
	// e.POST("/upload", m.handlePostUpload)
}

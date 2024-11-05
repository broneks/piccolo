package upload

import (
	"github.com/labstack/echo/v4"
)

func (m *UploadModule) Routes(e *echo.Echo) {
	e.GET("/uploads", m.handleGetUploads)
	e.POST("/upload", m.handlePostUpload)
}

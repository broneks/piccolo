package shared

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func isValidUUID(value string) bool {
	_, err := uuid.Parse(value)
	return err == nil
}

func GetIdParam(c echo.Context) string {
	id := c.Param("id")

	if id == "" || !isValidUUID(id) {
		return ""
	}

	return id
}

package api

import (
	"piccolo/api/upload"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {
	e := echo.New()

	e.Use(Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())

	e.Static("/", "public")

	upload.Router(e)

	e.HideBanner = true
	e.Logger.Fatal(e.Start(":8000"))
}

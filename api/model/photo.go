package model

import (
	"context"
	"fmt"
	"piccolo/api/storage/redis"
	"piccolo/api/storage/wasabi"
	"time"

	"github.com/labstack/echo/v4"
)

type Photo struct {
	Id          string
	Location    string
	Filename    string
	FileSize    int
	ContentType string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (p *Photo) GetUrl(c echo.Context) string {
	val, _ := redis.Get(c, p.Filename)
	if val == "" {
		fmt.Println("no cache found for url image.")
		url, _ := wasabi.GetPresignedUrl(context.Background(), p.Filename)
		redis.Set(c, p.Filename, url)
		return url
	} else {
		fmt.Println("found cached url for image")
		return val
	}
}

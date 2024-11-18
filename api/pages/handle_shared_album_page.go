package pages

import (
	"net/http"
	"piccolo/api/repo"
	"piccolo/api/util"
	"time"

	"github.com/labstack/echo/v4"
)

type Album struct {
	Name        string
	Description string
	CreatedAt   time.Time
}

type Test struct {
	Name string
}

type Payload struct {
	Album
	List []Test
}

func handleSharedAlbumPage(sharedAlbumRepo *repo.SharedAlbumRepo) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		albumId := util.GetIdParam(c)
		album, err := sharedAlbumRepo.GetById(ctx, albumId)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.Render(http.StatusOK, "album.html", &Payload{
			Album: Album{
				Name:        album.Name.String,
				Description: album.Name.String,
				CreatedAt:   album.CreatedAt.Time,
			},
			List: []Test{
				{Name: "foo"},
				{Name: "bar"},
				{Name: "baz"},
			},
		})
	}
}

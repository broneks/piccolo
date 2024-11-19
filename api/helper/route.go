package helper

import (
	"fmt"
	"log"
	"sort"
	"strings"

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

const staticRoute = "/static*"

func ListAllRoutes(e *echo.Echo) {
	routes := append([]*echo.Route(nil), e.Routes()...)

	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Path < routes[j].Path
	})

	fmt.Println()

	for _, route := range routes {
		if route.Method != "echo_route_not_found" && route.Path != staticRoute {
			method := route.Method

			switch method {
			case "GET":
				{
					method = method + strings.Repeat(" ", 3)
				}
			case "POST":
				{
					method = method + strings.Repeat(" ", 2)
				}
			}

			log.Printf("%s %s\n", method, route.Path)
		}
	}

	fmt.Println()
}

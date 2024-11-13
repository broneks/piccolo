package util

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/labstack/echo/v4"
)

const staticRoute = "/static*"

func ListAllRoutes(e *echo.Echo) {
	routes := append([]*echo.Route(nil), e.Routes()...)

	sort.Slice(routes, func(i, j int) bool {
		// force auth routes to list first
		if strings.Contains(routes[i].Path, "/auth/") {
			return true
		}
		if strings.Contains(routes[j].Path, "/auth/") {
			return false
		}

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

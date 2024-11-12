package util

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/labstack/echo/v4"
)

func ListAllRoutes(e *echo.Echo) {
	routes := append([]*echo.Route(nil), e.Routes()...)

	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Path < routes[j].Path
	})

	fmt.Println()

	for _, route := range routes {
		if route.Method != "echo_route_not_found" && route.Path != "/*" {
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

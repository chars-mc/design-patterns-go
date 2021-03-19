package main

import (
	"fmt"
	"os"

	"github.com/chars-mc/design-patterns-go/structural/decorator"
)

func main() {
	route := decorator.NewRoute()
	start(&route)

	var path string
	_, err := fmt.Scanf("%s", &path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	route.Exec(path)
}

func start(route *decorator.Route) {
	// decorating the route with LogRegistry
	route.Add(
		decorator.NewLogRegistry(decorator.NewHandlerSayHello()),
		"/greet",
	)
	route.Add(
		decorator.NewLogRegistry(decorator.NewHandlerSayGoodbye()),
		"/leave",
	)
}

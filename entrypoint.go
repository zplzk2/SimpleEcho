package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := flag.Int("port", 8080, "the port it listens on")
	flag.Parse()

	e := echo.New()

	e.Use(middleware.Recover())

	e.GET("/*", hello)

	e.Start(fmt.Sprintf(":%v", *port))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, c.Request().URL.String())
}

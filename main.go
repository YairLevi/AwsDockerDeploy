package main

import (
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	server := echo.New()

	server.GET("/", func(c echo.Context) error {
		return c.String(200, "Golang!")
	})

	server.GET("/what", func(c echo.Context) error {
		return c.String(200, "What?")
	})

	log.Fatal(server.Start(":8000"))
}

package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"main/handlers"
)

func main() {
	server := echo.New()
	h := handlers.Handler{}

	server.GET("/", h.Hello)
	server.GET("/what", h.Goodbye)

	log.Fatal(server.Start(":8000"))
}

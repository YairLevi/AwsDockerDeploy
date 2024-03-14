package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct{}

func (h *Handler) Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from handler 2")
}

func (h *Handler) Goodbye(c echo.Context) error {
	return c.String(http.StatusOK, "Goodbye from handler")
}

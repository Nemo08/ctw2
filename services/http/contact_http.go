package http

import (
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo/middleware"
)

type ContactHttpInterface struct {
	e *echo.Echo
}

func NewContactHttpInterface(e *echo.Echo) *ContactHttpInterface {
	return &ContactHttpInterface{
		e: e,
	}
}

func Create(c echo.Context) error {
	cont := &Contact{}
	if err := c.Bind(cont); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, cont)
}

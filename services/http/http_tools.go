package http

import (
	"context"
	_ "net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type HttpService struct {
	e        *echo.Echo
	basePath string
}

func (hs *HttpService) Start(port int) {
	// Start server
	go func() {
		if err := hs.e.Start(":" + strconv.Itoa(port)); err != nil {
			hs.e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := hs.e.Shutdown(ctx); err != nil {
		hs.e.Logger.Fatal(err)
	}
}

func (hs *HttpService) Add(method, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route {
	return hs.e.Add(method, path, handler, middleware...)
}

func NewHttpService(basepath string) *HttpService {
	e := echo.New()
	e.HideBanner = true //???wtf
	e.HidePort = true   //???wtf
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &HttpService{
		e:        echo.New(),
		basePath: basepath,
	}
}

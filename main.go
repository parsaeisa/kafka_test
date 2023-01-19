package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e *echo.Echo
}

func main() {
	s := &Server{}
	s.e = echo.New()

	s.e.GET("/", hello())

	if err := s.e.Start("localhost:8080"); err != nil && err != http.ErrServerClosed {
		s.e.Logger.Fatal("shutting down the server")
	}
}

func hello() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		return ctx.HTML(http.StatusOK, "<strong>Hello, World!</strong>")
	}
}

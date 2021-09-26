package server

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	e *echo.Echo
}

func New() *Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/new", newDB)
	e.GET("/:database", list)
	e.GET("/:database/:key", get)
	e.POST("/:database", set)
	e.PUT("/:database/:key", set)
	e.DELETE("/:database/:key", delete)
	return &Server{
		e: e,
	}
}

func (s *Server) Execute(port int) {
	s.e.Logger.Fatal(s.e.Start(":" + strconv.Itoa(port)))
}

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
	e.GET("/:dbname", list)
	e.GET("/:dbname/:key", get)
	e.POST("/:dbname", set)
	e.PUT("/:dbname/:key", set)
	e.DELETE("/:dbname/:key", delete)
	return &Server{
		e: e,
	}
}

func (s *Server) Execute(port int) {
	s.e.Logger.Fatal(s.e.Start(":" + strconv.Itoa(port)))
}

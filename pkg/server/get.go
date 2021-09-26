package server

import (
	"net/http"

	"github.com/delihiros/shockv/pkg/shockv"

	"github.com/labstack/echo/v4"
)

func get(c echo.Context) error {
	var req GetRequest
	if err := c.Bind(&req); err != nil || req.Database == "" || req.Key == "" {
		return c.JSON(http.StatusBadRequest, GetResponse{Response: &Response{Status: 400}})
	}
	db, err := shockv.Get()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, GetResponse{Response: &Response{Status: 500}})
	}
	bytes, err := db.Get(req.Database, req.Key)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, GetResponse{Response: &Response{Status: 500}})
	}
	return c.JSON(http.StatusOK, GetResponse{
		Response: &Response{Status: 200},
		Body:     bytes,
	})
}

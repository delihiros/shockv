package server

import (
	"net/http"

	"github.com/delihiros/shockv/pkg/shockv"

	"github.com/labstack/echo/v4"
)

func set(c echo.Context) error {
	var req SetRequest
	if err := c.Bind(&req); err != nil || req.DatabaseName == "" || req.Key == "" {
		return c.JSON(http.StatusBadRequest, SetResponse{Response: &Response{Status: 400}})
	}
	db, err := shockv.Get()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, SetResponse{Response: &Response{Status: 500}})
	}
	err = db.Set(req.DatabaseName, req.Key, req.Value)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, SetResponse{Response: &Response{Status: 500}})
	}
	return c.JSON(http.StatusCreated, SetResponse{
		Response: &Response{Status: 201},
	})
}

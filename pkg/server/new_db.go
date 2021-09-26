package server

import (
	"net/http"

	"github.com/delihiros/shockv/pkg/shockv"

	"github.com/labstack/echo/v4"
)

func newDB(c echo.Context) error {
	var req NewDBRequest
	if err := c.Bind(&req); err != nil || req.Database == "" {
		return c.JSON(http.StatusBadRequest, NewDBResponse{Response: &Response{Status: 400}})
	}
	db, err := shockv.Get()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewDBResponse{Response: &Response{Status: 500}})
	}
	if req.Diskless {
		err = db.NewDisklessDB(req.Database)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, NewDBResponse{Response: &Response{Status: 500}})
		}
	} else {
		err = db.NewDiskDB(req.Database)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, NewDBResponse{Response: &Response{Status: 500}})
		}
	}
	return c.JSON(http.StatusCreated, NewDBResponse{
		Response: &Response{Status: 201},
	})
}

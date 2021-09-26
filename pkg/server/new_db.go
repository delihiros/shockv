package server

import (
	"fmt"
	"net/http"

	"github.com/delihiros/shockv/pkg/protocols"

	"github.com/delihiros/shockv/pkg/shockv"

	"github.com/labstack/echo/v4"
)

func newDB(c echo.Context) error {
	var req protocols.NewDBRequest
	if err := c.Bind(&req); err != nil || req.Database == "" {
		return c.JSON(http.StatusBadRequest, protocols.NewDBResponse{
			Response: &protocols.Response{
				Status:  400,
				Message: fmt.Sprintf("err: %v", err),
			}})
	}
	db, err := shockv.Get()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, protocols.NewDBResponse{Response: &protocols.Response{Status: 500}})
	}
	if req.Diskless {
		err = db.NewDisklessDB(req.Database)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, protocols.NewDBResponse{Response: &protocols.Response{Status: 500}})
		}
	} else {
		err = db.NewDiskDB(req.Database)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, protocols.NewDBResponse{Response: &protocols.Response{Status: 500}})
		}
	}
	return c.JSON(http.StatusCreated, protocols.NewDBResponse{
		Response: &protocols.Response{Status: 201},
	})
}

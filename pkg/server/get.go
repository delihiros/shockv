package server

import (
	"fmt"
	"net/http"

	"github.com/delihiros/shockv/pkg/protocols"

	"github.com/delihiros/shockv/pkg/shockv"

	"github.com/labstack/echo/v4"
)

func get(c echo.Context) error {
	var req protocols.GetRequest
	if err := c.Bind(&req); err != nil || req.Database == "" || req.Key == "" {
		return c.JSON(http.StatusBadRequest, protocols.ListResponse{
			Response: &protocols.Response{
				Status:  400,
				Message: fmt.Sprintf("err: %v", err),
			}})
	}
	db, err := shockv.Get()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, protocols.GetResponse{Response: &protocols.Response{Status: 500}})
	}
	bytes, err := db.Get(req.Database, req.Key)
	if err != nil {
		return c.JSON(http.StatusNotFound, protocols.GetResponse{Response: &protocols.Response{Status: 404}})
	}
	return c.JSON(http.StatusOK, protocols.GetResponse{
		Response: &protocols.Response{Status: 200},
		Body:     bytes,
	})
}

package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/delihiros/shockv/pkg/protocols"

	"github.com/delihiros/shockv/pkg/shockv"

	"github.com/labstack/echo/v4"
)

func set(c echo.Context) error {
	var req protocols.SetRequest
	if err := c.Bind(&req); err != nil || req.Database == "" || req.Key == "" {
		return c.JSON(http.StatusBadRequest, protocols.ListResponse{
			Response: &protocols.Response{
				Status:  400,
				Message: fmt.Sprintf("err: %v", err),
			}})
	}
	ttl, err := strconv.Atoi(req.TTL)
	if err != nil {
		return c.JSON(http.StatusBadRequest, protocols.ListResponse{
			Response: &protocols.Response{
				Status:  400,
				Message: fmt.Sprintf("err: %v", fmt.Errorf("invalid TTL")),
			}})
	}
	db, err := shockv.Get()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, protocols.SetResponse{Response: &protocols.Response{Status: 500}})
	}
	err = db.Set(req.Database, req.Key, req.Value, ttl)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, protocols.SetResponse{Response: &protocols.Response{Status: 500}})
	}
	return c.JSON(http.StatusCreated, protocols.SetResponse{
		Response: &protocols.Response{Status: 201},
	})
}

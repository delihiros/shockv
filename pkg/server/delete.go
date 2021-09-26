package server

import (
	"fmt"
	"net/http"

	"github.com/delihiros/shockv/pkg/protocols"

	"github.com/delihiros/shockv/pkg/shockv"

	"github.com/labstack/echo/v4"
)

func delete(c echo.Context) error {
	var req protocols.DeleteRequest
	if err := c.Bind(&req); err != nil || req.Database == "" || req.Key == "" {
		return c.JSON(http.StatusBadRequest, protocols.ListResponse{
			Response: &protocols.Response{
				Status:  400,
				Message: fmt.Sprintf("err: %v", err),
			}})
	}
	db, err := shockv.Get()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, protocols.DeleteResponse{Response: &protocols.Response{Status: 500}})
	}
	err = db.Delete(req.Database, req.Key)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, protocols.DeleteResponse{Response: &protocols.Response{Status: 500}})
	}
	return c.JSON(http.StatusOK, protocols.DeleteResponse{
		Response: &protocols.Response{Status: http.StatusOK},
	})
}

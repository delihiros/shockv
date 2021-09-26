package server

import (
	"net/http"

	"github.com/delihiros/shockv/pkg/shockv"

	"github.com/labstack/echo/v4"
)

func delete(c echo.Context) error {
	var req DeleteRequest
	if err := c.Bind(&req); err != nil || req.Database == "" || req.Key == "" {
		return c.JSON(http.StatusBadRequest, DeleteResponse{Response: &Response{Status: 400}})
	}
	db, err := shockv.Get()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, DeleteResponse{Response: &Response{Status: 500}})
	}
	err = db.Delete(req.Database, req.Key)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, DeleteResponse{Response: &Response{Status: 500}})
	}
	return c.JSON(http.StatusNoContent, DeleteResponse{
		Response: &Response{Status: 204},
	})
}

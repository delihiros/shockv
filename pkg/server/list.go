package server

import (
	"net/http"

	"github.com/delihiros/shockv/pkg/shockv"

	"github.com/labstack/echo/v4"
)

func list(c echo.Context) error {
	var req ListRequest
	if err := c.Bind(&req); err != nil || req.Database == "" {
		return c.JSON(http.StatusBadRequest, ListResponse{Response: &Response{Status: 400}})
	}
	db, err := shockv.Get()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ListResponse{Response: &Response{Status: 500}})
	}
	kvMap, err := db.List(req.Database)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ListResponse{Response: &Response{Status: 500}})
	}
	pairs := []*Pair{}
	for k, v := range kvMap {
		pairs = append(pairs, &Pair{Key: k, Value: v})
	}
	return c.JSON(http.StatusOK, ListResponse{
		Response: &Response{Status: 200},
		Body:     pairs,
	})
}

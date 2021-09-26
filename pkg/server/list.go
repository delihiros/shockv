package server

import (
	"fmt"
	"net/http"

	"github.com/delihiros/shockv/pkg/protocols"

	"github.com/delihiros/shockv/pkg/shockv"

	"github.com/labstack/echo/v4"
)

func list(c echo.Context) error {
	var req protocols.ListRequest
	if err := c.Bind(&req); err != nil || req.Database == "" {
		return c.JSON(http.StatusBadRequest, protocols.ListResponse{
			Response: &protocols.Response{
				Status:  400,
				Message: fmt.Sprintf("err: %v", err),
			}})
	}
	db, err := shockv.Get()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, protocols.ListResponse{Response: &protocols.Response{Status: 500}})
	}
	kvMap, err := db.List(req.Database)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, protocols.ListResponse{Response: &protocols.Response{Status: 500}})
	}
	pairs := []*protocols.Pair{}
	for k, v := range kvMap {
		pairs = append(pairs, &protocols.Pair{Key: k, Value: v})
	}
	return c.JSON(http.StatusOK, protocols.ListResponse{
		Response: &protocols.Response{Status: 200},
		Body:     pairs,
	})
}

package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/nattrio/go-echo-demo/cmd/api/service"
)

func PostIndexHandler(c echo.Context) error {
	data, err := service.GetAll()
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process request")
	}

	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data

	return c.JSON(http.StatusOK, res)
}

func PostSingleHandler(c echo.Context) error {
	id := c.Param("id")
	idx, err := strconv.Atoi(id) // convert string to int
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process request")
	}

	data, err := service.GetById(idx)
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process request")
	}

	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data

	return c.JSON(http.StatusOK, res)
}

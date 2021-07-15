package app

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/infrastructure"
)

func GetController(c echo.Context) error {
	hours := c.QueryParam("hours")
	minutes := c.QueryParam("minutes")
	if response, err := infrastructure.EuToUsTimeConverter(hours, minutes); err == nil {
		return c.String(http.StatusOK, response)
	} else {
		return c.String(http.StatusBadRequest, err.Error())
	}
}
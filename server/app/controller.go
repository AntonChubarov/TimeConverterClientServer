package app

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/infrastructure"
)

func GetController(c echo.Context) error {
	hours := c.QueryParam("hours")
	minutes := c.QueryParam("minutes")
	return c.String(http.StatusOK, infrastructure.EuToUsTimeConverter(hours, minutes))
}
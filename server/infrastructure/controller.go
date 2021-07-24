package infrastructure

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/app"
)

func GetController(c echo.Context) error {
	hours := c.QueryParam("hours")
	minutes := c.QueryParam("minutes")
	if formattedTime, err := app.ConvertEuTimeToUs(hours, minutes); err == nil {
		return c.String(http.StatusOK, formattedTime)
	} else {
		return c.String(http.StatusBadRequest, err.Error())
	}
}
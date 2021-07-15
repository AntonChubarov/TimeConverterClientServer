package main

import (
	"github.com/labstack/echo/v4"
	"server/app"
)

func main() {
	e := echo.New()
	e.GET("/time", app.GetController)
	e.Logger.Fatal(e.Start(":8080"))
}

package main

import (
	"github.com/labstack/echo/v4"
	"server/infrastructure"
)

func main() {
	e := echo.New()
	e.GET("/time", infrastructure.GetController)
	e.Logger.Fatal(e.Start(":8080"))
}

package main

import (
	"client/app"
	"client/infrastructure"
)

func main() {
	converter := infrastructure.NewWebConverter()
	ui := infrastructure.NewConsole()
	app.Run(converter, ui)
}

package main

import (
	"github.com/marceloagmelo/go-auth-web/logger"

	"github.com/marceloagmelo/go-auth-web/app"
)

func main() {
	app := &app.App{}
	app.Initialize()
	logger.Info.Println("Listen 8080...")
	app.Run(":8080")
}

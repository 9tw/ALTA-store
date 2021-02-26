package main

import (
	"project/config"
	"project/routes"
	"project/middlewares"
)

func main() {
	config.InitDB()
	e := routes.New()
	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}

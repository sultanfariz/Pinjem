package main

import (
	"Pinjem/config"
	"Pinjem/routes"
	"log"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.LoadEnv()
	config.InitDB()
	e := routes.InitRoutes()
	e.Pre(middleware.RemoveTrailingSlash())
	log.Println(e.Start(":8080"))
}

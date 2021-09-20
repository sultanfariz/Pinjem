package main

import (
	"Pinjem/config"
	"Pinjem/controllers"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

func main() {
	LoadEnv()
	config.InitDB()
	e := echo.New()
	e.POST("/api/v1/register", controllers.RegisterController)
	e.POST("/api/v1/login", controllers.LoginController)
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// 	// Skipper: func(c echo.Context) bool {
	// 	// 	if strings.HasPrefix(c.Request().Host, "localhost") {
	// 	// 		return true
	// 	// 	}
	// 	// 	return false
	// 	// },
	// }))
	log.Println(e.Start(":8080"))
}

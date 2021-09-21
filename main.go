package main

import (
	"Pinjem/config"
	"Pinjem/controllers"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	// jwt := middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte(os.Getenv("JWT_SECRET")),
	// 	Claims:     jwt.MapClaims{},
	// })
	e.Pre(middleware.RemoveTrailingSlash())
	jwt := middleware.JWT([]byte(os.Getenv("JWT_SECRET")))
	// jwt := middleware.JWT([]byte("gilemangmantepparah"))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status} ${time_rfc3339} ${latency_human}\n",
		// Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}, latency=${latency_human}\n",
		// Skipper: func(c echo.Context) bool {
		// 	if strings.HasPrefix(c.Request().Host, "localhost") {
		// 		return true
		// 	}
		// 	return false
		// },
	}))
	e.POST("/api/v1/register", controllers.RegisterController)
	e.POST("/api/v1/login", controllers.LoginController)
	e.GET("/api/v1/users", controllers.GetAllUsersController, jwt)
	e.GET("/api/v1/users/:userId", controllers.GetUserByIdController)
	log.Println(e.Start(":8080"))
}

package main

import (
	"Pinjem/app/config"
	"Pinjem/app/routes"
	_userUsecase "Pinjem/businesses/users"
	_authController "Pinjem/controllers/auth"
	_userDb "Pinjem/drivers/databases/users"
	"Pinjem/drivers/mysql"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	config.LoadEnv()

	// connect to db
	configDB := mysql.ConfigDB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}
	Conn := configDB.InitDB()
	Conn.AutoMigrate(&_userDb.Users{})

	timeoutContextEnv, _ := strconv.Atoi(os.Getenv("TIMEOUT_CONTEXT"))
	timeoutContext := time.Duration(timeoutContextEnv) * time.Second

	userUsecase := _userUsecase.NewUsecase(_userDb.NewUserRepository(Conn), timeoutContext)
	authController := _authController.NewAuthController(*userUsecase)

	// Routes
	e.Pre(middleware.RemoveTrailingSlash())
	routesInit := routes.ControllerList{
		AuthController: authController,
	}
	routesInit.InitRoutes(e)
	log.Println(e.Start(":8080"))
}

package main

import (
	"Pinjem/app/config"
	"Pinjem/app/routes"
	_bookUsecase "Pinjem/businesses/books"
	_userUsecase "Pinjem/businesses/users"
	_authController "Pinjem/controllers/auth"
	_bookController "Pinjem/controllers/books"
	_userController "Pinjem/controllers/users"
	_bookDb "Pinjem/drivers/databases/books"
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
	Conn.AutoMigrate(&_userDb.Users{}, &_bookDb.Books{})

	timeoutContextEnv, _ := strconv.Atoi(os.Getenv("TIMEOUT_CONTEXT"))
	timeoutContext := time.Duration(timeoutContextEnv) * time.Second

	userUsecase := _userUsecase.NewUsecase(_userDb.NewUserRepository(Conn), timeoutContext)
	bookUseCase := _bookUsecase.NewUsecase(_bookDb.NewBookRepository(Conn), timeoutContext)
	authController := _authController.NewAuthController(*userUsecase)
	userController := _userController.NewUserController(*userUsecase)
	bookController := _bookController.NewBookController(*bookUseCase)

	// Routes
	e.Pre(middleware.RemoveTrailingSlash())
	routesInit := routes.ControllerList{
		AuthController: authController,
		UserController: userController,
		BookController: bookController,
	}
	routesInit.InitRoutes(e)
	log.Println(e.Start(":8080"))
}

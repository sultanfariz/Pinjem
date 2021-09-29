package main

import (
	"Pinjem/app/config"
	"Pinjem/app/routes"
	_bookUsecase "Pinjem/businesses/books"
	_depositUsecase "Pinjem/businesses/deposits"
	_userUsecase "Pinjem/businesses/users"
	_authController "Pinjem/controllers/auth"
	_bookController "Pinjem/controllers/books"
	_depositController "Pinjem/controllers/deposits"
	_userController "Pinjem/controllers/users"
	_bookDb "Pinjem/drivers/databases/books"
	_depositDb "Pinjem/drivers/databases/deposits"
	_userDb "Pinjem/drivers/databases/users"
	postgres "Pinjem/drivers/postgresql"
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
	configDB := postgres.ConfigDB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}
	Conn := configDB.InitDB()
	Conn.Debug().AutoMigrate(&_userDb.Users{}, &_bookDb.Books{}, &_depositDb.Deposits{})

	timeoutContextEnv, _ := strconv.Atoi(os.Getenv("TIMEOUT_CONTEXT"))
	timeoutContext := time.Duration(timeoutContextEnv) * time.Second

	userUsecase := _userUsecase.NewUsecase(_userDb.NewUserRepository(Conn), timeoutContext)
	bookUseCase := _bookUsecase.NewUsecase(_bookDb.NewBookRepository(Conn), timeoutContext)
	depositUseCase := _depositUsecase.NewUsecase(_depositDb.NewDepositRepository(Conn), timeoutContext)
	authController := _authController.NewAuthController(*userUsecase, *depositUseCase)
	userController := _userController.NewUserController(*userUsecase)
	bookController := _bookController.NewBookController(*bookUseCase)
	depositController := _depositController.NewDepositController(*depositUseCase)

	// Routes
	e.Pre(middleware.RemoveTrailingSlash())
	routesInit := routes.ControllerList{
		AuthController:    authController,
		UserController:    userController,
		BookController:    bookController,
		DepositController: depositController,
	}
	routesInit.InitRoutes(e)
	log.Println(e.Start(":8080"))
}

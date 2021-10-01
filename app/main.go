package main

import (
	"Pinjem/app/config"
	"Pinjem/app/routes"
	_bookOrderUsecase "Pinjem/businesses/book_orders"
	_bookUsecase "Pinjem/businesses/books"
	_depositUsecase "Pinjem/businesses/deposits"
	_orderUsecase "Pinjem/businesses/orders"
	_userUsecase "Pinjem/businesses/users"
	_authController "Pinjem/controllers/auth"
	_bookOrderController "Pinjem/controllers/book_orders"
	_bookController "Pinjem/controllers/books"
	_depositController "Pinjem/controllers/deposits"
	_orderController "Pinjem/controllers/orders"
	_userController "Pinjem/controllers/users"
	_bookOrderDb "Pinjem/drivers/databases/book_orders"
	_bookDb "Pinjem/drivers/databases/books"
	_depositDb "Pinjem/drivers/databases/deposits"
	_orderDb "Pinjem/drivers/databases/orders"
	_userDb "Pinjem/drivers/databases/users"
	postgres "Pinjem/drivers/postgresql"

	// mysql "Pinjem/drivers/mysql"
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
		// configDB := mysql.ConfigDB{
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
	orderUseCase := _orderUsecase.NewUsecase(_orderDb.NewOrderRepository(Conn), timeoutContext)
	bookOrderUseCase := _bookOrderUsecase.NewUsecase(_bookOrderDb.NewBookOrderRepository(Conn), timeoutContext)
	authController := _authController.NewAuthController(*userUsecase, *depositUseCase)
	userController := _userController.NewUserController(*userUsecase, *depositUseCase)
	bookController := _bookController.NewBookController(*bookUseCase)
	depositController := _depositController.NewDepositController(*depositUseCase)
	orderController := _orderController.NewOrderController(*orderUseCase, *bookOrderUseCase)
	bookOrderController := _bookOrderController.NewBookOrderController(*bookOrderUseCase)

	// Routes
	e.Pre(middleware.RemoveTrailingSlash())
	routesInit := routes.ControllerList{
		AuthController:      authController,
		UserController:      userController,
		BookController:      bookController,
		DepositController:   depositController,
		OrderController:     orderController,
		BookOrderController: bookOrderController,
	}
	routesInit.InitRoutes(e)
	log.Println(e.Start(":8080"))
}

package routes

import (
	"Pinjem/controllers/auth"
	bookOrders "Pinjem/controllers/book_orders"
	"Pinjem/controllers/books"
	"Pinjem/controllers/deposits"
	"Pinjem/controllers/orders"
	"Pinjem/controllers/users"
	"Pinjem/helpers"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	AuthController      *auth.AuthController
	UserController      *users.UserController
	BookController      *books.BookController
	DepositController   *deposits.DepositController
	OrderController     *orders.OrderController
	BookOrderController *bookOrders.BookOrderController
}

func (c ControllerList) InitRoutes(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	// v1.Use(middleware.Recover())
	v1.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	jwt := middleware.JWT([]byte(os.Getenv("JWT_SECRET")))
	v1.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status} ${time_rfc3339} ${latency_human}\n",
	}))

	// unprotected routes
	{
		v1.Static("/uploads", "public")

		// auth routes
		v1.POST("/register", c.AuthController.Register)
		v1.POST("/login", c.AuthController.Login)

		// book routes
		v1.GET("/books", c.BookController.GetAll)
		v1.GET("/books/:bookId", c.BookController.GetById)
		// v1.POST("/orders/my", c.OrderController.Create, jwt)
	}

	// admin routes
	admins := v1.Group("")
	admins.Use(helpers.AdminRoleValidation)
	{
		// user routes
		admins.GET("/users", c.UserController.GetAll, jwt)
		admins.GET("/users/:userId", c.UserController.GetById, jwt)

		// book routes
		admins.POST("/books", c.BookController.Create, jwt)

		// deposit routes
		admins.GET("/deposits/:userId", c.DepositController.GetByUserId, jwt)
		admins.GET("/deposits", c.DepositController.GetAll, jwt)

		// // order routes
		// admins.GET("/orders", c.OrderController.GetAll, jwt)
		// admins.GET("/orders/:orderId", c.OrderController.GetById, jwt)
		// // admins.GET("/:userId", c.OrderController.GetByUserId, jwt)
	}

	// user routes
	users := v1.Group("")
	users.Use(helpers.UserRoleValidation)
	{
		// user routes
		users.GET("/users/my", c.UserController.GetMyUserProfile, jwt)

		// deposit routes
		users.POST("/deposits/my", c.DepositController.Update, jwt)

		// order routes
		users.POST("/orders", c.OrderController.Create, jwt)
		// users.GET("/my", c.OrderController.GetMyOrders, jwt)
	}

	// // user routes
	// users := v1.Group("/users")
	// users.GET("/my", c.UserController.GetMyUserProfile, jwt)

	// users.Use(helpers.AdminRoleValidation)
	// users.GET("/all", c.UserController.GetAll, jwt)
	// users.GET("/:userId", c.UserController.GetById, jwt)

	// // book routes
	// books := v1.Group("/books")
	// {
	// 	books.GET("/:bookId", c.BookController.GetById)
	// 	books.GET("/all", c.BookController.GetAll)
	// }
	// adminBooks := v1.Group("/books")
	// adminBooks.Use(helpers.AdminRoleValidation)
	// {
	// 	adminBooks.POST("", c.BookController.Create, jwt)
	// }

	// // deposit routes
	// userDeposits := v1.Group("/deposits")
	// userDeposits.Use(helpers.UserRoleValidation)
	// {
	// 	userDeposits.POST("/my", c.DepositController.Update, jwt)
	// }
	// adminDeposits := v1.Group("/deposits")
	// adminDeposits.Use(helpers.AdminRoleValidation)
	// {
	// 	adminDeposits.GET("/:userId", c.DepositController.GetByUserId, jwt)
	// 	adminDeposits.GET("", c.DepositController.GetAll, jwt)
	// }

	// // order routes
	// userOrders := v1.Group("/orders")
	// userOrders.Use(helpers.UserRoleValidation)
	// {
	// 	userOrders.POST("", c.OrderController.Create, jwt)
	// 	// userOrders.GET("/my", c.OrderController.GetMyOrders, jwt)
	// }
	// adminOrders := v1.Group("/orders")
	// adminOrders.Use(helpers.AdminRoleValidation)
	// {
	// 	adminOrders.GET("", c.OrderController.GetAll, jwt)
	// 	adminOrders.GET("/:orderId", c.OrderController.GetById, jwt)
	// 	// adminOrders.GET("/:userId", c.OrderController.GetByUserId, jwt)
	// }

	// bookOrder routes

}

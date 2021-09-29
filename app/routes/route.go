package routes

import (
	"Pinjem/controllers/auth"
	"Pinjem/controllers/books"
	"Pinjem/controllers/deposits"
	"Pinjem/controllers/users"
	"Pinjem/helpers"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	AuthController    *auth.AuthController
	UserController    *users.UserController
	BookController    *books.BookController
	DepositController *deposits.DepositController
}

func (c ControllerList) InitRoutes(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	v1.Use(middleware.Recover())
	v1.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	jwt := middleware.JWT([]byte(os.Getenv("JWT_SECRET")))
	// userRoleValidation := helpers.UserRoleValidation(v1, v1)
	v1.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status} ${time_rfc3339} ${latency_human}\n",
	}))

	v1.Static("/uploads", "public")

	// auth routes
	v1.POST("/register", c.AuthController.Register)
	v1.POST("/login", c.AuthController.Login)

	// user routes
	users := v1.Group("/users")
	users.GET("/my", c.UserController.GetMyUserProfile, jwt)

	users.Use(helpers.AdminRoleValidation)
	users.GET("/all", c.UserController.GetAll, jwt)
	users.GET("/:userId", c.UserController.GetById, jwt)

	// book routes
	// v1.GET("/books", c.BookController.GetAll)
	// books.Use(helpers.UserRoleValidation)
	books := v1.Group("/books")
	books.GET("/:bookId", c.BookController.GetById)
	books.GET("/all", c.BookController.GetAll)
	books.Use(helpers.AdminRoleValidation)
	books.POST("", c.BookController.Create, jwt)
	// books.POST("/books/:isbn", c.BookController.Create, jwt)
	// books.POST("/books", c.BookController.Create, jwt, helpers.UserRoleValidation)
	// books.POST("/books/:userId", c.BookController.Create)
	// books.PUT("/books/:bookId", c.BookController.Update, jwt)

	// deposit routes
	deposits := v1.Group("/deposits")
	deposits.POST("/my", c.DepositController.Update, jwt)
	deposits.Use(helpers.AdminRoleValidation)
	deposits.GET("/:userId", c.DepositController.GetByUserId, jwt)
	deposits.GET("", c.DepositController.GetAll, jwt)

}

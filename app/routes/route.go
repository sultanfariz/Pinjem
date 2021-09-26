package routes

import (
	"Pinjem/controllers/auth"
	"Pinjem/controllers/books"
	"Pinjem/controllers/users"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	AuthController *auth.AuthController
	UserController *users.UserController
	BookController *books.BookController
}

func (c ControllerList) InitRoutes(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	v1.Use(middleware.Recover())
	v1.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	jwt := middleware.JWT([]byte(os.Getenv("JWT_SECRET")))
	v1.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status} ${time_rfc3339} ${latency_human}\n",
	}))

	v1.Static("/uploads", "public")

	// user routes
	v1.POST("/register", c.AuthController.Register)
	v1.POST("/login", c.AuthController.Login)
	v1.GET("/users", c.UserController.GetAll, jwt)
	v1.GET("/users/:userId", c.UserController.GetById)

	// book routes
	// v1.GET("/books", c.BookController.GetAll, jwt)
	v1.GET("/books/:bookId", c.BookController.GetById)
	v1.POST("/books/:userId", c.BookController.Create, jwt)
	// v1.POST("/books/:userId", c.BookController.Create)
	// v1.PUT("/books/:bookId", c.BookController.Update, jwt)
}

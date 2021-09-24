package routes

import (
	"Pinjem/controllers/auth"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	AuthController *auth.AuthController
	// UserController *controllers.UserController
	// BookController *controllers.BookController
}

func (c ControllerList) InitRoutes(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	// v1.Use(middleware.Logger())
	v1.Use(middleware.Recover())
	v1.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// jwt := middleware.JWT([]byte(os.Getenv("JWT_SECRET")))
	v1.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status} ${time_rfc3339} ${latency_human}\n",
	}))

	v1.Static("/uploads", "public")

	// v1.POST("/register", c.AuthController.Register)
	v1.POST("/login", c.AuthController.Login)
	// v1.GET("/users", c.UserController.GetAllUsers, jwt)
	// v1.GET("/users/:userId", c.UserController.GetUserById, jwt)
}

// func InitRoutes() *echo.Echo {
// 	e := echo.New()
// 	v1 := e.Group("/api/v1")
// 	// jwt := middleware.JWTWithConfig(middleware.JWTConfig{
// 	// 	SigningKey: []byte(os.Getenv("JWT_SECRET")),
// 	// 	Claims:     jwt.MapClaims{},
// 	// })
// 	jwt := middleware.JWT([]byte(os.Getenv("JWT_SECRET")))
// 	v1.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
// 		Format: "${method} ${uri} ${status} ${time_rfc3339} ${latency_human}\n",
// 		// Skipper: func(c echo.Context) bool {
// 		// 	if strings.HasPrefix(c.Request().Host, "localhost") {
// 		// 		return true
// 		// 	}
// 		// 	return false
// 		// },
// 	}))
// 	v1.Use(middleware.Recover())
// 	v1.Static("/uploads", "public")

// 	v1.POST("/register", controllers.RegisterController)
// 	v1.POST("/login", controllers.LoginController)
// 	v1.GET("/users", controllers.GetAllUsersController, jwt)
// 	v1.GET("/users/:userId", controllers.GetUserByIdController)

// 	return e
// }

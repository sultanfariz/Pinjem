package routes

import (
	"Pinjem/controllers"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes() *echo.Echo {
	e := echo.New()
	v1 := e.Group("/api/v1")
	// jwt := middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte(os.Getenv("JWT_SECRET")),
	// 	Claims:     jwt.MapClaims{},
	// })
	jwt := middleware.JWT([]byte(os.Getenv("JWT_SECRET")))
	v1.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status} ${time_rfc3339} ${latency_human}\n",
		// Skipper: func(c echo.Context) bool {
		// 	if strings.HasPrefix(c.Request().Host, "localhost") {
		// 		return true
		// 	}
		// 	return false
		// },
	}))

	v1.POST("/register", controllers.RegisterController)
	v1.POST("/login", controllers.LoginController)
	v1.GET("/users", controllers.GetAllUsersController, jwt)
	v1.GET("/users/:userId", controllers.GetUserByIdController)

	return e
}

package routes

import (
	"Pinjem/controllers/auth"
	bookOrders "Pinjem/controllers/book_orders"
	"Pinjem/controllers/books"
	"Pinjem/controllers/deposits"
	"Pinjem/controllers/orders"
	"Pinjem/controllers/users"
	"Pinjem/helpers"
	"html/template"
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
	// render unit test result
	renderer := &helpers.TemplateRenderer{
		Templates: template.Must(template.ParseGlob("*.html")),
	}
	e.Renderer = renderer
	e.GET("/api/coverage", func(c echo.Context) error {
		return c.Render(200, "cover.html", nil)
	})

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
		admins.GET("/orders", c.OrderController.GetAll, jwt)
		admins.GET("/orders/:orderId", c.OrderController.GetById, jwt)
		admins.PATCH("/orders/:orderId", c.OrderController.UpdateStatus, jwt)
		admins.DELETE("/orders/:orderId", c.OrderController.Delete, jwt)
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
		users.GET("/orders/my", c.OrderController.GetMyOrders, jwt)
	}
}

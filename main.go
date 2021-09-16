package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct {
	Name     string
	Email    string
	Password string
}

type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

func GetUser(c echo.Context) error {
	user := User{Name: "Alterra", Email: "alterra@alterra.id"}
	return c.JSON(http.StatusOK, user)
}

func PostRegisterController(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	var response Response
	var user User

	if name == "" || email == "" || password == "" {
		response.Status = http.StatusBadRequest
		response.Success = false
		response.Message = "Please fill all the fields"
		response.Content = ""
	} else {
		user.Name = name
		user.Email = email
		user.Password = password

		response.Status = http.StatusOK
		response.Success = true
		response.Message = "Register success"
		response.Content = user
	}

	return c.JSON(response.Status, response)
}

func main() {
	e := echo.New()
	e.GET("/api/v1/user", GetUser)
	e.POST("/api/v1/register", PostRegisterController)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	log.Println(e.Start(":8080"))
}

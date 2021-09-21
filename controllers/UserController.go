package controllers

import (
	"Pinjem/config"
	users "Pinjem/models/user"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

func GetAllUsersController(c echo.Context) error {
	users := []users.User{}
	result := config.DB.Find(&users)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, Response{Status: http.StatusInternalServerError, Success: false, Message: "Internal Server Error"})
		}
	}

	return c.JSON(http.StatusOK, Response{Status: http.StatusOK, Success: true, Message: "OK", Content: users})
}

func GetUserByIdController(c echo.Context) error {
	users := users.User{}
	result := config.DB.First(&users, c.Param("userId"))
	fmt.Println(result.Error)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, Response{Status: http.StatusInternalServerError, Success: false, Message: "Internal Server Error"})
		}
	}

	return c.JSON(http.StatusOK, Response{Status: http.StatusOK, Success: true, Message: "OK", Content: users})
}

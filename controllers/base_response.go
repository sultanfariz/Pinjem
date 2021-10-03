package controllers

import (
	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c echo.Context, data interface{}) error {
	res := BaseResponse{
		Status:  200,
		Success: true,
		Message: "Success",
		Data:    data,
	}
	return c.JSON(res.Status, res)
}

func ErrorResponse(c echo.Context, status int, err error) error {
	res := BaseResponse{
		Status:  status,
		Success: false,
		Message: err.Error(),
		Data:    nil,
	}
	return c.JSON(res.Status, res)
}

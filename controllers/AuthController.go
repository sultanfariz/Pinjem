package controllers

import (
	"Pinjem/config"
	"Pinjem/helpers"
	users "Pinjem/models/user"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

func LoginController(c echo.Context) error {
	var userLogin users.LoginUserinput
	var response Response
	c.Bind(&userLogin)
	if userLogin.Email == "" || userLogin.Password == "" {
		response.Status = http.StatusBadRequest
		response.Success = false
		response.Message = "Please fill all the fields"
		response.Content = ""
		return c.JSON(response.Status, response)
	}
	user := config.DB.Where(&users.User{Email: userLogin.Email}).First(&users.User{})
	if user.Error != nil {
		fmt.Println(user)
		response.Status = http.StatusBadRequest
		response.Success = false
		response.Message = "error pas ngecek ke db"
		response.Content = ""
		return c.JSON(response.Status, response)
	}

	return nil
}

func RegisterController(c echo.Context) error {
	var user users.User
	var userRegister users.RegisterUserinput
	var response Response
	var err error
	c.Bind(&userRegister)

	if userRegister.Fullname == "" || userRegister.Email == "" || userRegister.Password == "" {
		response.Status = http.StatusBadRequest
		response.Success = false
		response.Message = "Please fill all the fields"
		response.Content = ""
	} else {
		user.Email = userRegister.Email
		user.Password, err = helpers.HashPassword(userRegister.Password)
		if err != nil {
			response.Status = http.StatusInternalServerError
			response.Success = false
			response.Message = "Internal Server Error"
			response.Content = ""
		} else {
			user.Fullname = userRegister.Fullname
			user.NIK = userRegister.NIK
			user.PhoneNumber = userRegister.PhoneNumber
			user.Birthdate = userRegister.Birthdate
			user.Address = userRegister.Address
			user.Provinsi = userRegister.Provinsi
			user.Kota = userRegister.Kota
			user.Kecamatan = userRegister.Kecamatan
			user.Kelurahan = userRegister.Kelurahan
			user.PostalCode = userRegister.PostalCode
			user.Role = userRegister.Role
			user.Status = 1
			user.CreatedAt = time.Now()
			user.UpdatedAt = time.Now()
			res := config.DB.Where(&users.User{Email: user.Email}).FirstOrCreate(&user)
			if res.Error != nil {
				response.Status = http.StatusInternalServerError
				response.Success = false
				response.Message = "Internal Server Error"
				response.Content = ""
			} else {
				response.Status = http.StatusOK
				response.Success = true
				response.Message = "Register success"
				response.Content = user
			}
		}
	}

	return c.JSON(response.Status, response)
}

package controllers

import (
	"Pinjem/config"
	"Pinjem/helpers"
	responses "Pinjem/models/response"
	users "Pinjem/models/user"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	var userLogin users.LoginUserinput
	var err error
	userLogin.Email = c.FormValue("email")
	userLogin.Password = c.FormValue("password")
	// c.Bind(&userLogin)
	// fmt.Println(userLogin)
	if userLogin.Email == "" || userLogin.Password == "" {
		return c.JSON(http.StatusBadRequest, responses.Response{
			Status:  http.StatusBadRequest,
			Success: false,
			Message: "Please fill all the fields",
		})
	}
	// search email from db
	user := users.User{}
	result := config.DB.Where("email = ?", userLogin.Email).Find(&user)
	if result.Error != nil || user.ID == 0 {
		return c.JSON(http.StatusInternalServerError, responses.Response{
			Status:  http.StatusUnauthorized,
			Success: false,
			Message: "Username or password doesn't match our records",
		})
	}

	// check password
	res := helpers.IsMatched(userLogin.Password, user.Password)
	if !res {
		return c.JSON(http.StatusUnauthorized, responses.Response{
			Status:  http.StatusUnauthorized,
			Success: false,
			Message: "Username or password doesn't match our records",
		})
	}
	// generate jwt token
	token, err := helpers.GenerateToken(int(user.ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.Response{
			Status:  http.StatusInternalServerError,
			Success: false,
			Message: "Internal Server Error",
		})
	}
	return c.JSON(http.StatusOK, responses.Response{
		Status:  http.StatusOK,
		Success: true,
		Message: "Login success",
		Content: struct {
			Token string `json:"token"`
		}{token},
	})
}

func RegisterController(c echo.Context) error {
	var user users.User
	var userRegister users.RegisterUserinput
	// c.Bind(&userRegister)
	userRegister.Email = c.FormValue("email")
	userRegister.Password = c.FormValue("password")
	userRegister.Fullname = c.FormValue("fullname")
	userRegister.NIK = c.FormValue("nik")
	userRegister.PhoneNumber = c.FormValue("phoneNumber")
	userRegister.Birthdate = c.FormValue("birthdate")
	userRegister.Address = c.FormValue("address")
	userRegister.Provinsi = c.FormValue("provinsi")
	userRegister.Kota = c.FormValue("kota")
	userRegister.Kecamatan = c.FormValue("kecamatan")
	userRegister.Desa = c.FormValue("desa")
	userRegister.PostalCode = c.FormValue("postalCode")
	userRegister.Role = c.FormValue("role")

	if userRegister.Fullname == "" || userRegister.Email == "" || userRegister.Password == "" {
		return c.JSON(http.StatusBadRequest, responses.Response{
			Status:  http.StatusBadRequest,
			Success: false,
			Message: "Please fill all the fields",
		})
	}
	hash, err := helpers.HashPassword(userRegister.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.Response{
			Status:  http.StatusInternalServerError,
			Success: false,
			Message: "Internal Server Error",
		})
	}
	user.Email = userRegister.Email
	user.Password = hash
	user.Fullname = userRegister.Fullname
	user.NIK = userRegister.NIK
	user.PhoneNumber = userRegister.PhoneNumber
	user.Birthdate = userRegister.Birthdate
	user.Address = userRegister.Address
	user.Provinsi = userRegister.Provinsi
	user.Kota = userRegister.Kota
	user.Kecamatan = userRegister.Kecamatan
	user.Desa = userRegister.Desa
	user.PostalCode = userRegister.PostalCode
	user.Role = userRegister.Role
	user.Status = 1
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	// insert user to db
	res := config.DB.Where(&users.User{Email: user.Email}).FirstOrCreate(&user)
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, responses.Response{
			Status:  http.StatusInternalServerError,
			Success: false,
			Message: "Internal Server Error",
		})
	}
	return c.JSON(http.StatusOK, responses.Response{
		Status:  http.StatusOK,
		Success: true,
		Message: "User registered successfully",
		Content: user,
	})
}

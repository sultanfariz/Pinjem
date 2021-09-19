package controllers

import (
	"Pinjem/config"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type User struct {
	ID          uint   `gorm:"primary_key"`
	Email       string `gorm:"unique_index" json: "email"`
	Password    string `gorm:"not null" json: "password"`
	Fullname    string `gorm:"not null" json: "fullname"`
	NIK         string `gorm:"not null" json: "nik"`
	PhoneNumber string `gorm:"not null" json: "phoneNumber"`
	Birthdate   string `gorm:"not null" json: "birthdate"`
	Address     string `gorm:"not null" json: "address"`
	Provinsi    string `gorm:"not null" json: "provinsi"`
	Kota        string `gorm:"not null" json: "kota"`
	Kecamatan   string `gorm:"not null" json: "kecamatan"`
	Kelurahan   string `gorm:"not null" json: "kelurahan"`
	PostalCode  string `gorm:"not null" json: "postalCode"`
	Role        string `gorm:"not null" json: "role"`
	Status      int    `gorm:"not null" json: "status"`
	CreatedAt   time.Time `json: "createdAt"`
	UpdatedAt   time.Time `json: "updatedAt"`
}

type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

func RegisterController(c echo.Context) error {
	var user User
	var response Response
	c.Bind(&user)

	if user.Fullname == "" || user.Email == "" || user.Password == "" {
		response.Status = http.StatusBadRequest
		response.Success = false
		response.Message = "Please fill all the fields"
		response.Content = ""
	} else {
		config.DB.Where(&User{Email: user.Email}).FirstOrCreate(&user)
		response.Status = http.StatusOK
		response.Success = true
		response.Message = "Register success"
		response.Content = user
	}

	return c.JSON(response.Status, response)
}

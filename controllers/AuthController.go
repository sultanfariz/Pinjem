package controllers

import (
	"Pinjem/config"
	"Pinjem/helpers"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User struct {
	ID          uint           `gorm:"primary_key" json:"id"`
	Email       string         `gorm:"unique_index" json:"email"`
	Password    string         `gorm:"not null" json:"password"`
	Fullname    string         `gorm:"not null" json:"fullname"`
	NIK         string         `gorm:"not null" json:"nik"`
	PhoneNumber string         `gorm:"not null" json:"phone_number"`
	Birthdate   string         `gorm:"not null" json:"birthdate"`
	Address     string         `gorm:"not null" json:"address"`
	Provinsi    string         `gorm:"not null" json:"provinsi"`
	Kota        string         `gorm:"not null" json:"kota"`
	Kecamatan   string         `gorm:"not null" json:"kecamatan"`
	Kelurahan   string         `gorm:"not null" json:"kelurahan"`
	PostalCode  string         `gorm:"not null" json:"postalCode"`
	Role        string         `gorm:"not null" json:"role"`
	Status      int            `gorm:"not null" json:"status"`
	CreatedAt   time.Time      `json: "createdAt"`
	UpdatedAt   time.Time      `json: "updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
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
	var err error
	// c.Bind(&user)
	user.Email = c.FormValue("email")
	user.Password, err = helpers.HashPassword(c.FormValue("password"))
	user.Fullname = c.FormValue("fullname")
	user.NIK = c.FormValue("nik")
	user.PhoneNumber = c.FormValue("phoneNumber")
	user.Birthdate = c.FormValue("birthdate")
	user.Address = c.FormValue("address")
	user.Provinsi = c.FormValue("provinsi")
	user.Kota = c.FormValue("kota")
	user.Kecamatan = c.FormValue("kecamatan")
	user.Kelurahan = c.FormValue("kelurahan")
	user.PostalCode = c.FormValue("postalCode")
	user.Role = c.FormValue("role")
	user.Status = 1
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.DeletedAt = gorm.DeletedAt{}

	if user.Fullname == "" || user.Email == "" || user.Password == "" {
		response.Status = http.StatusBadRequest
		response.Success = false
		response.Message = "Please fill all the fields"
		response.Content = ""
	} else {
		// user.Password, err = hashPassword(user.Password)
		if err != nil {
			response.Status = http.StatusInternalServerError
			response.Success = false
			response.Message = "Internal Server Error"
			response.Content = ""
		} else {
			// user.Password, err = helpers.hashPassword(user.Password)
			fmt.Println(user)
			fmt.Println(user.Password)
			res := config.DB.Where(&User{Email: user.Email}).FirstOrCreate(&user)
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

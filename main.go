package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// database config
var (
	DB *gorm.DB
)

func InitDB() {
	dsn := "root:@/pinjem?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Migration()
}

type User struct {
	ID          uint   `gorm:"primary_key"`
	Email       string `gorm:"unique_index"`
	Password    string `gorm:"not null"`
	Fullname    string `gorm:"not null"`
	NIK         string `gorm:"not null"`
	PhoneNumber string `gorm:"not null"`
	Birthdate   string `gorm:"not null"`
	Address     string `gorm:"not null"`
	Provinsi    string `gorm:"not null"`
	Kota        string `gorm:"not null"`
	Kecamatan   string `gorm:"not null"`
	Kelurahan   string `gorm:"not null"`
	PostalCode  string `gorm:"not null"`
	Role        string `gorm:"not null"`
	Status      int    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func Migration() {
	DB.AutoMigrate(&User{})
}

type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

func GetUser(c echo.Context) error {
	user := User{Fullname: "Alterra", Email: "alterra@alterra.id"}
	return c.JSON(http.StatusOK, user)
}

func PostRegisterController(c echo.Context) error {
	var user User
	var response Response
	c.Bind(&user)

	if user.Fullname == "" || user.Email == "" || user.Password == "" {
		response.Status = http.StatusBadRequest
		response.Success = false
		response.Message = "Please fill all the fields"
		response.Content = ""
	} else {
		DB.Where(&User{Email: user.Email}).FirstOrCreate(&user)
		response.Status = http.StatusOK
		response.Success = true
		response.Message = "Register success"
		response.Content = user
	}

	return c.JSON(response.Status, response)
}

func main() {
	InitDB()
	e := echo.New()
	e.GET("/api/v1/user", GetUser)
	e.POST("/api/v1/register", PostRegisterController)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	log.Println(e.Start(":8080"))
}

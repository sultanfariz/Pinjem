package main

import (
	"Pinjem/config"
	"Pinjem/controllers"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

type User struct {
	ID          uint      `gorm:"primary_key"`
	Email       string    `gorm:"unique_index" json: "email"`
	Password    string    `gorm:"not null" json: "password"`
	Fullname    string    `gorm:"not null" json: "fullname"`
	NIK         string    `gorm:"not null" json: "nik"`
	PhoneNumber string    `gorm:"not null" json: "phoneNumber"`
	Birthdate   string    `gorm:"not null" json: "birthdate"`
	Address     string    `gorm:"not null" json: "address"`
	Provinsi    string    `gorm:"not null" json: "provinsi"`
	Kota        string    `gorm:"not null" json: "kota"`
	Kecamatan   string    `gorm:"not null" json: "kecamatan"`
	Kelurahan   string    `gorm:"not null" json: "kelurahan"`
	PostalCode  string    `gorm:"not null" json: "postalCode"`
	Role        string    `gorm:"not null" json: "role"`
	Status      int       `gorm:"not null" json: "status"`
	CreatedAt   time.Time `json: "createdAt"`
	UpdatedAt   time.Time `json: "updatedAt"`
}

type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

func main() {
	LoadEnv()
	config.InitDB()
	e := echo.New()
	e.POST("/api/v1/register", controllers.RegisterController)
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// 	// Skipper: func(c echo.Context) bool {
	// 	// 	if strings.HasPrefix(c.Request().Host, "localhost") {
	// 	// 		return true
	// 	// 	}
	// 	// 	return false
	// 	// },
	// }))
	log.Println(e.Start(":8080"))
}

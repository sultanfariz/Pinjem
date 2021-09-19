package config

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

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

func Migration() {
	DB.AutoMigrate(&User{})
}

func InitDB() {
	dsn, exists := os.LookupEnv("DSN")
	var err error
	if !exists {
		log.Fatal("DSN not defined in .env file")
	}
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Migration()
}

package config

import (
	users "Pinjem/models/user"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Migration() {
	DB.AutoMigrate(&users.User{})
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

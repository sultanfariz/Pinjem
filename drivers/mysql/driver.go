package mysql

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func (c *ConfigDB) InitDB() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return db
}

// var (
// 	DB *gorm.DB
// )

// func Migration() {
// 	DB.AutoMigrate(&users.User{})
// }

// func InitDB() {
// 	dsn, exists := os.LookupEnv("DSN")
// 	var err error
// 	if !exists {
// 		log.Fatal("DSN not defined in .env file")
// 	}
// 	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	Migration()
// }

package postgresql

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
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
	dsn := "host=localhost user=postgres password=postgresql dbname=pinjem port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return db
}

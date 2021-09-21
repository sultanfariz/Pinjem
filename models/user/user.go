package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `gorm:"primary_key"`
	Email       string         `gorm:"unique_index" json:"email"`
	Password    string         `gorm:"not null" json:"password"`
	Fullname    string         `gorm:"not null" json:"fullname"`
	NIK         string         `gorm:"not null" json:"nik"`
	PhoneNumber string         `gorm:"not null" json:"phoneNumber"`
	Birthdate   string         `gorm:"not null" json:"birthdate"`
	Address     string         `gorm:"not null" json:"address"`
	Provinsi    string         `gorm:"not null" json:"provinsi"`
	Kota        string         `gorm:"not null" json:"kota"`
	Kecamatan   string         `gorm:"not null" json:"kecamatan"`
	Desa        string         `gorm:"not null" json:"desa"`
	PostalCode  string         `gorm:"not null" json:"postalCode"`
	Role        string         `gorm:"not null" json:"role"`
	Status      int            `gorm:"not null" json:"status"`
	Token       string         `json:"token" form:"token"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type RegisterUserinput struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Fullname    string `json:"fullname"`
	NIK         string `json:"nik"`
	PhoneNumber string `json:"phoneNumber"`
	Birthdate   string `json:"birthdate"`
	Address     string `json:"address"`
	Provinsi    string `json:"provinsi"`
	Kota        string `json:"kota"`
	Kecamatan   string `json:"kecamatan"`
	Desa        string `json:"desa"`
	PostalCode  string `json:"postalCode"`
	Role        string `json:"role"`
	Status      int    `json:"status"`
}

type LoginUserinput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

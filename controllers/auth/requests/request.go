package requests

import (
	"time"

	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	ID          uint
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	Fullname    string         `json:"fullname"`
	NIK         string         `json:"nik"`
	PhoneNumber string         `json:"phoneNumber"`
	Birthdate   string         `json:"birthdate"`
	Address     string         `json:"address"`
	Provinsi    string         `json:"provinsi"`
	Kota        string         `json:"kota"`
	Kecamatan   string         `json:"kecamatan"`
	Desa        string         `json:"desa"`
	PostalCode  string         `json:"postalCode"`
	Role        string         `json:"role"`
	Status      int            `json:"status"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
}

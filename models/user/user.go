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
	Role        string         `gorm:"not null;type:enum('admin', 'customer');default:'customer'" json:"role"`
	Status      int            `gorm:"not null" json:"status"`
	LinkKTP     string         `gorm:"not null" json:"linkKTP"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type RegisterUserinput struct {
	Email       string `json:"email" form:"email" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
	Fullname    string `json:"fullname" form:"fullname" binding:"required"`
	NIK         string `json:"nik" form:"nik" binding:"required"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber" binding:"required"`
	Birthdate   string `json:"birthdate" form:"birthdate" binding:"required"`
	Address     string `json:"address" form:"address" binding:"required"`
	Provinsi    string `json:"provinsi" form:"provinsi" binding:"required"`
	Kota        string `json:"kota" form:"kota" binding:"required"`
	Kecamatan   string `json:"kecamatan" form:"kecamatan" binding:"required"`
	Desa        string `json:"desa" form:"desa" binding:"required"`
	PostalCode  string `json:"postalCode" form:"postalCode" binding:"required"`
	Role        string `json:"role" form:"role" binding:"required"`
	Status      int    `json:"status" form:"status" binding:"required"`
	LinkKTP     string `json:"linkKTP" form:"linkKTP" binding:"required"`
}

type LoginUserinput struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

package responses

import (
	"Pinjem/businesses/users"
	"time"
)

type UserResponse struct {
	ID          uint
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Fullname    string    `json:"fullname"`
	NIK         string    `json:"nik"`
	PhoneNumber string    `json:"phoneNumber"`
	Birthdate   string    `json:"birthdate"`
	Address     string    `json:"address"`
	Provinsi    string    `json:"provinsi"`
	Kota        string    `json:"kota"`
	Kecamatan   string    `json:"kecamatan"`
	Desa        string    `json:"desa"`
	PostalCode  string    `json:"postalCode"`
	Role        string    `json:"role"`
	LinkKTP     string    `json:"linkKTP"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	ID          uint
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Fullname    string    `json:"fullname"`
	NIK         string    `json:"nik"`
	PhoneNumber string    `json:"phoneNumber"`
	Birthdate   string    `json:"birthdate"`
	Address     string    `json:"address"`
	Provinsi    string    `json:"provinsi"`
	Kota        string    `json:"kota"`
	Kecamatan   string    `json:"kecamatan"`
	Desa        string    `json:"desa"`
	PostalCode  string    `json:"postalCode"`
	Role        string    `json:"role"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		ID:          domain.Id,
		Email:       domain.Email,
		Password:    domain.Password,
		Fullname:    domain.Fullname,
		NIK:         domain.Nik,
		PhoneNumber: domain.PhoneNumber,
		Birthdate:   domain.Birthdate,
		Address:     domain.Address,
		Provinsi:    domain.Provinsi,
		Kota:        domain.Kota,
		Kecamatan:   domain.Kecamatan,
		Desa:        domain.Desa,
		PostalCode:  domain.PostalCode,
		Role:        domain.Role,
		Status:      domain.Status,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

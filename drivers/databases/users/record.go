package users

import (
	"Pinjem/businesses/users"
	"time"

	"gorm.io/gorm"
)

type Users struct {
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
	Desa        string `gorm:"not null"`
	PostalCode  string `gorm:"not null"`
	// Role        string `gorm:"not null;type:enum('admin', 'customer');default:'customer'"`
	Role      string `gorm:"not null;default:'customer'"`
	Status    int    `gorm:"not null"`
	LinkKTP   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Users) TableName() string {
	return "users"
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

func (u *Users) ToDomain() users.Domain {
	return users.Domain{
		Id:          u.ID,
		Email:       u.Email,
		Password:    u.Password,
		Fullname:    u.Fullname,
		Nik:         u.NIK,
		PhoneNumber: u.PhoneNumber,
		Birthdate:   u.Birthdate,
		Address:     u.Address,
		Provinsi:    u.Provinsi,
		Kota:        u.Kota,
		Kecamatan:   u.Kecamatan,
		Desa:        u.Desa,
		PostalCode:  u.PostalCode,
		Role:        u.Role,
		Status:      u.Status,
		LinkKTP:     u.LinkKTP,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

func FromDomain(domain users.Domain) Users {
	return Users{
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
		LinkKTP:     domain.LinkKTP,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func ToListDomain(data []Users) []users.Domain {
	var listDomain []users.Domain
	for _, d := range data {
		listDomain = append(listDomain, d.ToDomain())
	}
	return listDomain
}

package users

import (
	context "context"
	"time"
)

type Domain struct {
	Id          uint
	Fullname    string
	Email       string
	Password    string
	Nik         string
	PhoneNumber string
	Address     string
	Birthdate   string
	Provinsi    string
	Kota        string
	Kecamatan   string
	Desa        string
	PostalCode  string
	Role        string
	Status      int
	LinkKTP     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type DomainRepository interface {
	// GetAll(ctx context.Context) ([]Domain, error)
	// GetById(ctx context.Context, id int64) (Domain, error)
	// Create(ctx context.Context, domain Domain) (Domain, error)
	// Update(ctx context.Context, domain Domain) (Domain, error)
	// Delete(ctx context.Context, id int64) (Domain, error)
	Login(ctx context.Context, email string, password string) (Domain, error)
}

type DomainService interface {
	// GetAll(ctx context.Context) ([]Domain, error)
	// GetById(ctx context.Context, id int64) (Domain, error)
	// Create(ctx context.Context, domain Domain) (Domain, error)
	// Update(ctx context.Context, domain Domain) (Domain, error)
	// Delete(ctx context.Context, id int64) (Domain, error)
	Login(ctx context.Context, domain Domain) (Domain, error)
	// Login(ctx context.Context, email string, password string) (Domain, error)
}

// 	ID          uint           `gorm:"primary_key"`
// 	Email       string         `gorm:"unique_index" json:"email"`
// 	Password    string         `gorm:"not null" json:"password"`
// 	Fullname    string         `gorm:"not null" json:"fullname"`
// 	NIK         string         `gorm:"not null" json:"nik"`
// 	PhoneNumber string         `gorm:"not null" json:"phoneNumber"`
// 	Birthdate   string         `gorm:"not null" json:"birthdate"`
// 	Address     string         `gorm:"not null" json:"address"`
// 	Provinsi    string         `gorm:"not null" json:"provinsi"`
// 	Kota        string         `gorm:"not null" json:"kota"`
// 	Kecamatan   string         `gorm:"not null" json:"kecamatan"`
// 	Desa        string         `gorm:"not null" json:"desa"`
// 	PostalCode  string         `gorm:"not null" json:"postalCode"`
// 	Role        string         `gorm:"not null;type:enum('admin', 'customer');default:'customer'" json:"role"`
// 	Status      int            `gorm:"not null" json:"status"`
// 	LinkKTP     string         `gorm:"not null" json:"linkKTP"`
// 	CreatedAt   time.Time      `json:"createdAt"`
// 	UpdatedAt   time.Time      `json:"updatedAt"`
// 	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
// }

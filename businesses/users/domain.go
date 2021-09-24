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
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	// Update(ctx context.Context, domain Domain) (Domain, error)
	// Delete(ctx context.Context, id uint) (Domain, error)
	Login(ctx context.Context, email string, password string) (Domain, error)
	// Register(ctx context.Context, domain Domain) (Domain, error)
	FindByEmail(ctx context.Context, email string) (Domain, error)
}

type DomainService interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	// Update(ctx context.Context, domain Domain) (Domain, error)
	// Delete(ctx context.Context, id uint) (Domain, error)
	Login(ctx context.Context, domain Domain) (Domain, error)
	FindByEmail(ctx context.Context, email string) (Domain, error)
}

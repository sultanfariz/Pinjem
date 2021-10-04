package users

import (
	context "context"
	"time"
)

type Domain struct {
	Id          uint
	Fullname    string `validate:"required"`
	Email       string `validate:"required,email"`
	Password    string `validate:"required,min=8,max=50"`
	Nik         string `validate:"required,len=16"`
	PhoneNumber string `validate:"required,min=10"`
	Address     string `validate:"required"`
	Birthdate   string `validate:"required"`
	Provinsi    string `validate:"required"`
	Kota        string `validate:"required"`
	Kecamatan   string `validate:"required"`
	Desa        string `validate:"required"`
	PostalCode  string `validate:"required"`
	Role        string `validate:"required"`
	Status      int    `validate:"required"`
	LinkKTP     string `validate:"required,url"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type LoginDomain struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=50"`
}

type DomainRepository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	// Update(ctx context.Context, domain Domain) (Domain, error)
	// Delete(ctx context.Context, id uint) (Domain, error)
	Login(ctx context.Context, email string, password string) (Domain, error)
	FindByEmail(ctx context.Context, email string) (Domain, error)
}

type DomainService interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Register(ctx context.Context, domain Domain) (Domain, error)
	// Update(ctx context.Context, domain Domain) (Domain, error)
	// Delete(ctx context.Context, id uint) (Domain, error)
	Login(ctx context.Context, email string, password string) (Domain, error)
	FindByEmail(ctx context.Context, email string) (Domain, error)
}

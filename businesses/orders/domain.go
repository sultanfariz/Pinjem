package orders

import (
	context "context"
	"time"
)

type Domain struct {
	Id        uint
	UserId    uint
	OrderDate time.Time
	ExpDate   time.Time
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type DomainRepository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetOrdersByUserId(ctx context.Context, userId uint) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	// Update(ctx context.Context, domain Domain) (Domain, error)
	Delete(ctx context.Context, id uint) error
}

type DomainService interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetOrdersByUserId(ctx context.Context, userId uint) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	// Update(ctx context.Context, domain Domain) (Domain, error)
	Delete(ctx context.Context, id uint) error
}

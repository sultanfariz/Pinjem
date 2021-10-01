package book_orders

import (
	context "context"
	"time"
)

type Domain struct {
	Id            uint
	OrderId       uint
	BookId        string
	DepositAmount uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type DomainRepository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id string) (Domain, error)
	GetByOrderId(ctx context.Context, id uint) ([]Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	// Update(ctx context.Context, domain Domain) (Domain, error)
	// Delete(ctx context.Context, id uint) (Domain, error)
}

type DomainService interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id string) (Domain, error)
	GetByOrderId(ctx context.Context, id uint) ([]Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	// Update(ctx context.Context, domain Domain) (Domain, error)
	// Delete(ctx context.Context, id uint) (Domain, error)
}

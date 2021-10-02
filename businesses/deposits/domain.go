package deposits

import (
	context "context"
	"time"
)

type Domain struct {
	Id         uint
	UserId     uint
	Amount     uint
	UsedAmount uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

type DomainRepository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByUserId(ctx context.Context, userId uint) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, userId uint, amount uint, usedAmount uint) (Domain, error)
	TopUp(ctx context.Context, userId uint, amount uint) (Domain, error)
	// Delete(ctx context.Context, id uint) (Domain, error)
}

type DomainService interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByUserId(ctx context.Context, userId uint) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, userId uint, amount uint, usedAmount uint) (Domain, error)
	TopUp(ctx context.Context, userId uint, amount uint) (Domain, error)
	// Delete(ctx context.Context, id uint) (Domain, error)
}

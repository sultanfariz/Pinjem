package shipping_details

import (
	context "context"
	"time"
)

type Domain struct {
	Id             uint
	OrderId        uint
	DestProvinsi   string
	DestKota       string
	DestKecamatan  string
	DestDesa       string
	DestAddress    string
	DestPostalCode string
	ShippingCost   uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

type DomainRepository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByOrderId(ctx context.Context, orderId uint) (Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	Delete(ctx context.Context, id uint) error
	DeleteByOrderId(ctx context.Context, orderId uint) error
}

type DomainService interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByOrderId(ctx context.Context, orderId uint) (Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	Delete(ctx context.Context, id uint) error
	DeleteByOrderId(ctx context.Context, orderId uint) error
}

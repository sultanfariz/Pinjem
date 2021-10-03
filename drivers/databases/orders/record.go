package orders

import (
	"Pinjem/businesses/orders"
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	ID        uint `gorm:"primary_key"`
	UserId    uint `gorm:"not null"`
	OrderDate time.Time
	ExpDate   time.Time
	Status    bool `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (o *Orders) BeforeCreate() (err error) {
	o.OrderDate = time.Now()
	o.ExpDate = time.Now().AddDate(0, 0, 60)
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
	return
}

func (o *Orders) ToDomain() orders.Domain {
	return orders.Domain{
		Id:        o.ID,
		UserId:    o.UserId,
		OrderDate: o.OrderDate,
		ExpDate:   o.ExpDate,
		Status:    o.Status,
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}
}

func FromDomain(domain orders.Domain) Orders {
	return Orders{
		ID:        domain.Id,
		UserId:    domain.UserId,
		OrderDate: domain.OrderDate,
		ExpDate:   domain.ExpDate,
		Status:    domain.Status,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ToListDomain(data []Orders) []orders.Domain {
	var listDomain []orders.Domain
	for _, d := range data {
		listDomain = append(listDomain, d.ToDomain())
	}
	return listDomain
}

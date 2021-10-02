package shipping_details

import (
	shippingDetails "Pinjem/businesses/shipping_details"
	"time"

	"gorm.io/gorm"
)

type ShippingDetails struct {
	ID             uint   `gorm:"primary_key"`
	OrderId        uint   `gorm:"not null"`
	DestProvinsi   string `gorm:"not null"`
	DestKota       string `gorm:"not null"`
	DestKecamatan  string `gorm:"not null"`
	DestDesa       string `gorm:"not null"`
	DestAddress    string `gorm:"not null"`
	DestPostalCode string `gorm:"not null"`
	ShippingCost   uint   `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (s *ShippingDetails) BeforeCreate() (err error) {
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
	return
}

func (s *ShippingDetails) ToDomain() shippingDetails.Domain {
	return shippingDetails.Domain{
		Id:             s.ID,
		OrderId:        s.OrderId,
		DestProvinsi:   s.DestProvinsi,
		DestKota:       s.DestKota,
		DestKecamatan:  s.DestKecamatan,
		DestDesa:       s.DestDesa,
		DestAddress:    s.DestAddress,
		DestPostalCode: s.DestPostalCode,
		ShippingCost:   s.ShippingCost,
		CreatedAt:      s.CreatedAt,
		UpdatedAt:      s.UpdatedAt,
	}
}

func FromDomain(domain shippingDetails.Domain) ShippingDetails {
	return ShippingDetails{
		ID:            domain.Id,
		OrderId:       domain.OrderId,
		DestProvinsi:  domain.DestProvinsi,
		DestKota:      domain.DestKota,
		DestKecamatan: domain.DestKecamatan,
		DestDesa:      domain.DestDesa,
		DestAddress:   domain.DestAddress,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}

func ToListDomain(data []ShippingDetails) []shippingDetails.Domain {
	var listDomain []shippingDetails.Domain
	for _, d := range data {
		listDomain = append(listDomain, d.ToDomain())
	}
	return listDomain
}

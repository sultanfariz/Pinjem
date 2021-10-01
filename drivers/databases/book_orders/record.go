package book_orders

import (
	bookOrders "Pinjem/businesses/book_orders"
	"time"

	"gorm.io/gorm"
)

type BookOrders struct {
	ID            uint   `gorm:"primary_key"`
	OrderId       uint   `gorm:"column:order_id;not null"`
	BookId        string `gorm:"column:book_id;not null"`
	DepositAmount uint   `gorm:"column:deposit_amount;not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (b *BookOrders) BeforeCreate(tx *gorm.DB) (err error) {
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	return
}

func (b *BookOrders) ToDomain() bookOrders.Domain {
	return bookOrders.Domain{
		Id:            b.ID,
		OrderId:       b.OrderId,
		BookId:        b.BookId,
		DepositAmount: b.DepositAmount,
		CreatedAt:     b.CreatedAt,
		UpdatedAt:     b.UpdatedAt,
	}
}

func FromDomain(domain bookOrders.Domain) BookOrders {
	return BookOrders{
		ID:            domain.Id,
		OrderId:       domain.OrderId,
		BookId:        domain.BookId,
		DepositAmount: domain.DepositAmount,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}

func ToListDomain(data []BookOrders) []bookOrders.Domain {
	var listDomain []bookOrders.Domain
	for _, d := range data {
		listDomain = append(listDomain, d.ToDomain())
	}
	return listDomain
}

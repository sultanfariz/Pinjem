package book_orders

import (
	bookOrders "Pinjem/businesses/book_orders"
	"context"
	"time"

	"gorm.io/gorm"
)

type BookOrderRepository struct {
	Conn *gorm.DB
}

func NewBookOrderRepository(conn *gorm.DB) bookOrders.DomainRepository {
	return &BookOrderRepository{Conn: conn}
}

func (b *BookOrderRepository) GetAll(ctx context.Context) ([]bookOrders.Domain, error) {
	var bookOrdersModel []BookOrders
	if err := b.Conn.Find(&bookOrdersModel).Error; err != nil {
		return nil, err
	}
	var result []bookOrders.Domain
	result = ToListDomain(bookOrdersModel)
	return result, nil
}

func (b *BookOrderRepository) GetById(ctx context.Context, id string) (bookOrders.Domain, error) {
	var book BookOrders
	if err := b.Conn.Where("book_id = ?", id).First(&book).Error; err != nil {
		return bookOrders.Domain{}, err
	}
	return book.ToDomain(), nil
}

func (b *BookOrderRepository) GetByOrderId(ctx context.Context, id uint) ([]bookOrders.Domain, error) {
	var bookOrdersModel []BookOrders
	if err := b.Conn.Where("order_id = ?", id).Find(&bookOrdersModel).Error; err != nil {
		return []bookOrders.Domain{}, err
	}
	var result []bookOrders.Domain
	result = ToListDomain(bookOrdersModel)
	return result, nil
}

func (b *BookOrderRepository) Create(ctx context.Context, bookOrder bookOrders.Domain) (bookOrders.Domain, error) {
	createdBookOrder := BookOrders{
		OrderId:       bookOrder.OrderId,
		BookId:        bookOrder.BookId,
		DepositAmount: bookOrder.DepositAmount,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	insertErr := b.Conn.Create(&createdBookOrder).Error
	if insertErr != nil {
		return bookOrders.Domain{}, insertErr
	}
	return createdBookOrder.ToDomain(), nil
}

func (b *BookOrderRepository) Delete(ctx context.Context, id string) error {
	var book BookOrders
	if err := b.Conn.Where("id = ?", id).Delete(&book).Error; err != nil {
		return err
	}
	return nil
}

func (b *BookOrderRepository) DeleteByOrderId(ctx context.Context, orderId uint) error {
	var book BookOrders
	if err := b.Conn.Where("order_id = ?", orderId).Delete(&book).Error; err != nil {
		return err
	}
	return nil
}

// func (b *BookOrderRepository) Update(user *User) error {
// 	return b.Conn.Save(user).Error
// }

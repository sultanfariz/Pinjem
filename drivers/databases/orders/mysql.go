package orders

import (
	"Pinjem/businesses/orders"
	"context"
	"log"

	"gorm.io/gorm"
)

type OrderRepository struct {
	Conn *gorm.DB
}

func NewOrderRepository(conn *gorm.DB) orders.DomainRepository {
	return &OrderRepository{Conn: conn}
}

func (b *OrderRepository) GetAll(ctx context.Context) ([]orders.Domain, error) {
	var ordersModel []Orders
	if err := b.Conn.Find(&ordersModel).Error; err != nil {
		return nil, err
	}
	var result []orders.Domain
	result = ToListDomain(ordersModel)
	return result, nil
}

func (b *OrderRepository) GetById(ctx context.Context, id string) (orders.Domain, error) {
	var order Orders
	if err := b.Conn.Where("order_id = ?", id).First(&order).Error; err != nil {
		return orders.Domain{}, err
	}
	return order.ToDomain(), nil
}

func (b *OrderRepository) GetByISBN(ctx context.Context, isbn string) (orders.Domain, error) {
	var order Orders
	if err := b.Conn.Where("isbn = ?", isbn).First(&order).Error; err != nil {
		return orders.Domain{}, err
	}
	return order.ToDomain(), nil
}

func (b *OrderRepository) Create(ctx context.Context, order orders.Domain) (orders.Domain, error) {
	createdOrder := Orders{
		UserId: order.UserId,
		Status: order.Status,
	}
	createdOrder.BeforeCreate()
	log.Println(createdOrder)

	insertErr := b.Conn.Create(&createdOrder).Error
	if insertErr != nil {
		return orders.Domain{}, insertErr
	}
	return createdOrder.ToDomain(), nil
}

// func (b *OrderRepository) Update(user *User) error {
// 	return b.Conn.Save(user).Error
// }

// func (b *OrderRepository) Delete(user *User) error {
// 	return b.Conn.Delete(user).Error
// }

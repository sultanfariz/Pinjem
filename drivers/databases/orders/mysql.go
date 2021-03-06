package orders

import (
	"Pinjem/businesses/orders"
	"context"
	"time"

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

func (b *OrderRepository) GetOrdersByUserId(ctx context.Context, userId uint) ([]orders.Domain, error) {
	var ordersModel []Orders
	if err := b.Conn.Where("user_id = ?", userId).Find(&ordersModel).Error; err != nil {
		return nil, err
	}
	var result []orders.Domain
	result = ToListDomain(ordersModel)
	return result, nil
}

func (b *OrderRepository) GetById(ctx context.Context, id uint) (orders.Domain, error) {
	var order Orders
	if err := b.Conn.Where("id = ?", id).First(&order).Error; err != nil {
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

	insertErr := b.Conn.Create(&createdOrder).Error
	if insertErr != nil {
		return orders.Domain{}, insertErr
	}
	return createdOrder.ToDomain(), nil
}

func (b *OrderRepository) UpdateStatus(ctx context.Context, id uint, status bool) (orders.Domain, error) {
	var orderModel Orders
	if err := b.Conn.Where("id = ?", id).First(&orderModel).Error; err != nil {
		return orders.Domain{}, err
	}
	orderModel.Status = status
	orderModel.UpdatedAt = time.Now()
	if err := b.Conn.Save(&orderModel).Error; err != nil {
		return orders.Domain{}, err
	}
	return orderModel.ToDomain(), nil
}

func (b *OrderRepository) Delete(ctx context.Context, id uint) error {
	var order Orders
	if err := b.Conn.Where("id = ?", id).Delete(&order).Error; err != nil {
		return err
	}
	return nil
}

// func (b *OrderRepository) Update(user *User) error {
// 	return b.Conn.Save(user).Error
// }

package shipping_details

import (
	shippingDetails "Pinjem/businesses/shipping_details"
	"context"

	"gorm.io/gorm"
)

type ShippingDetailsRepository struct {
	Conn *gorm.DB
}

func NewShippingDetailsRepository(conn *gorm.DB) shippingDetails.DomainRepository {
	return &ShippingDetailsRepository{Conn: conn}
}

func (b *ShippingDetailsRepository) GetAll(ctx context.Context) ([]shippingDetails.Domain, error) {
	var shippingDetailsModel []ShippingDetails
	if err := b.Conn.Find(&shippingDetailsModel).Error; err != nil {
		return nil, err
	}
	var result []shippingDetails.Domain
	result = ToListDomain(shippingDetailsModel)
	return result, nil
}

func (b *ShippingDetailsRepository) GetByOrderId(ctx context.Context, userId uint) ([]shippingDetails.Domain, error) {
	var shippingDetailsModel []ShippingDetails
	if err := b.Conn.Where("user_id = ?", userId).Find(&shippingDetailsModel).Error; err != nil {
		return nil, err
	}
	var result []shippingDetails.Domain
	result = ToListDomain(shippingDetailsModel)
	return result, nil
}

func (b *ShippingDetailsRepository) GetById(ctx context.Context, id uint) (shippingDetails.Domain, error) {
	var order ShippingDetails
	if err := b.Conn.Where("id = ?", id).First(&order).Error; err != nil {
		return shippingDetails.Domain{}, err
	}
	return order.ToDomain(), nil
}

func (b *ShippingDetailsRepository) Create(ctx context.Context, order shippingDetails.Domain) (shippingDetails.Domain, error) {
	createdShippingDetail := FromDomain(order)
	createdShippingDetail.BeforeCreate()

	insertErr := b.Conn.Create(&createdShippingDetail).Error
	if insertErr != nil {
		return shippingDetails.Domain{}, insertErr
	}
	return createdShippingDetail.ToDomain(), nil
}

func (b *ShippingDetailsRepository) Delete(ctx context.Context, id uint) error {
	var order ShippingDetails
	if err := b.Conn.Where("id = ?", id).Delete(&order).Error; err != nil {
		return err
	}
	return nil
}

// func (b *ShippingDetailsRepository) Update(user *User) error {
// 	return b.Conn.Save(user).Error
// }

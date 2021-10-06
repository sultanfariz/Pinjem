package shipping_details

import (
	shippingDetails "Pinjem/businesses/shipping_details"
	"context"

	"gorm.io/gorm"
)

type ShippingDetailRepository struct {
	Conn *gorm.DB
}

func NewShippingDetailRepository(conn *gorm.DB) shippingDetails.DomainRepository {
	return &ShippingDetailRepository{Conn: conn}
}

func (b *ShippingDetailRepository) GetAll(ctx context.Context) ([]shippingDetails.Domain, error) {
	var shippingDetailsModel []ShippingDetails
	if err := b.Conn.Find(&shippingDetailsModel).Error; err != nil {
		return nil, err
	}
	var result []shippingDetails.Domain
	result = ToListDomain(shippingDetailsModel)
	return result, nil
}

func (b *ShippingDetailRepository) GetByOrderId(ctx context.Context, orderId uint) (shippingDetails.Domain, error) {
	var shippingDetail ShippingDetails
	if err := b.Conn.Where("order_id = ?", orderId).Find(&shippingDetail).Error; err != nil {
		return shippingDetails.Domain{}, err
	}
	var result shippingDetails.Domain
	result = shippingDetail.ToDomain()
	return result, nil
}

func (b *ShippingDetailRepository) GetById(ctx context.Context, id uint) (shippingDetails.Domain, error) {
	var shippingDetail ShippingDetails
	if err := b.Conn.Where("id = ?", id).First(&shippingDetail).Error; err != nil {
		return shippingDetails.Domain{}, err
	}
	return shippingDetail.ToDomain(), nil
}

func (b *ShippingDetailRepository) Create(ctx context.Context, shippingDetail shippingDetails.Domain) (shippingDetails.Domain, error) {
	createdShippingDetail := FromDomain(shippingDetail)
	createdShippingDetail.BeforeCreate()

	err := b.Conn.Create(&createdShippingDetail).Error
	if err != nil {
		return shippingDetails.Domain{}, err
	}
	return createdShippingDetail.ToDomain(), nil
}

func (b *ShippingDetailRepository) Delete(ctx context.Context, id uint) error {
	var shippingDetail ShippingDetails
	if err := b.Conn.Where("id = ?", id).Delete(&shippingDetail).Error; err != nil {
		return err
	}
	return nil
}

func (b *ShippingDetailRepository) DeleteByOrderId(ctx context.Context, orderId uint) error {
	var order ShippingDetails
	if err := b.Conn.Where("order_id = ?", orderId).Delete(&order).Error; err != nil {
		return err
	}
	return nil
}

// func (b *ShippingDetailRepository) Update(user *User) error {
// 	return b.Conn.Save(user).Error
// }

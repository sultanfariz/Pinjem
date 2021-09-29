package deposits

import (
	"Pinjem/businesses/deposits"
	"context"
	"time"

	"gorm.io/gorm"
)

type DepositRepository struct {
	Conn *gorm.DB
}

func NewDepositRepository(conn *gorm.DB) deposits.DomainRepository {
	return &DepositRepository{Conn: conn}
}

func (d *DepositRepository) GetAll(ctx context.Context) ([]deposits.Domain, error) {
	var depositsModel []Deposits
	if err := d.Conn.Find(&depositsModel).Error; err != nil {
		return nil, err
	}
	var result []deposits.Domain
	result = ToListDomain(depositsModel)
	return result, nil
}

func (d *DepositRepository) GetByUserId(ctx context.Context, userId uint) (deposits.Domain, error) {
	var deposit Deposits
	if err := d.Conn.Where("user_id = ?", userId).First(&deposit).Error; err != nil {
		return deposits.Domain{}, err
	}
	return deposit.ToDomain(), nil
}

func (d *DepositRepository) Create(ctx context.Context, deposit deposits.Domain) (deposits.Domain, error) {
	createdDeposit := Deposits{
		UserID:    deposit.UserId,
		Amount:    deposit.Amount,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	insertErr := d.Conn.Create(&createdDeposit).Error
	if insertErr != nil {
		return deposits.Domain{}, insertErr
	}
	return createdDeposit.ToDomain(), nil
}

func (d *DepositRepository) Update(ctx context.Context, userId uint, amount uint) (deposits.Domain, error) {
	var deposit Deposits
	if err := d.Conn.Where("user_id = ?", userId).First(&deposit).Error; err != nil {
		return deposits.Domain{}, err
	}
	deposit.Amount += amount
	deposit.UpdatedAt = time.Now()
	if err := d.Conn.Save(&deposit).Error; err != nil {
		return deposits.Domain{}, err
	}
	return deposit.ToDomain(), nil
}

// func (d *DepositRepository) Delete(deposit *Deposit) error {
// 	return u.Conn.Delete(deposit).Error
// }

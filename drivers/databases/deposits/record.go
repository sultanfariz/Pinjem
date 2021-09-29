package deposits

import (
	deposits "Pinjem/businesses/deposits"
	"Pinjem/drivers/databases/users"
	"time"

	"gorm.io/gorm"
)

type Deposits struct {
	ID        uint        `gorm:"primary_key"`
	UserID    uint        `gorm:"column:user_id"`
	User      users.Users `gorm:"foreignkey:UserID"`
	Amount    uint        `gorm:"column:amount"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Deposits) TableName() string {
	return "deposits"
}

func (d *Deposits) BeforeCreate(tx *gorm.DB) (err error) {
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
	return
}

func (d *Deposits) ToDomain() deposits.Domain {
	return deposits.Domain{
		Id:        d.ID,
		UserId:    d.UserID,
		Amount:    d.Amount,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func FromDomain(domain deposits.Domain) Deposits {
	return Deposits{
		ID:        domain.Id,
		UserID:    domain.UserId,
		Amount:    domain.Amount,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ToListDomain(data []Deposits) []deposits.Domain {
	var listDomain []deposits.Domain
	for _, d := range data {
		listDomain = append(listDomain, d.ToDomain())
	}
	return listDomain
}

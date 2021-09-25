package books

import (
	"Pinjem/businesses/books"
	"time"

	"gorm.io/gorm"
)

type Books struct {
	ID          uint   `gorm:"primary_key"`
	PublisherId string `gorm:"type:varchar(100);not null"`
	ISBN        string `gorm:"type:varchar(13);not null"`
	Title       string `gorm:"type:varchar(100);not null"`
	// Category	string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:text;not null"`
	MinDeposit  uint   `gorm:"not null"`
	Status      bool   `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (u *Books) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

func (u *Books) ToDomain() books.Domain {
	return books.Domain{
		Id:          u.ID,
		PublisherId: u.PublisherId,
		ISBN:        u.ISBN,
		Title:       u.Title,
		// Category:    u.Category,
		Description: u.Description,
		MinDeposit:  u.MinDeposit,
		Status:      u.Status,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

func FromDomain(domain books.Domain) Books {
	return Books{
		ID:          domain.Id,
		PublisherId: domain.PublisherId,
		ISBN:        domain.ISBN,
		Title:       domain.Title,
		// Category:    domain.Category,
		Description: domain.Description,
		MinDeposit:  domain.MinDeposit,
		Status:      domain.Status,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func ToListDomain(data []Books) []books.Domain {
	var listDomain []books.Domain
	for _, d := range data {
		listDomain = append(listDomain, d.ToDomain())
	}
	return listDomain
}

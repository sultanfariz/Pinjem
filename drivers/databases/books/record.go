package books

import (
	"Pinjem/businesses/books"
	"time"

	"gorm.io/gorm"
)

// type Books struct {
// 	ID          uint     `gorm:"primary_key"`
// 	ISBN        string   `gorm:"type:varchar(13);not null"`
// 	Publisher   []string `gorm:"type:varchar(100);not null"`
// 	PublishDate string   `gorm:"type:varchar(100);not null"`
// 	Title       string   `gorm:"type:varchar(100);not null"`
// 	// Category	string `gorm:"type:varchar(100);not null"`
// 	Description   string `gorm:"type:text;not null"`
// 	MinDeposit    uint   `gorm:"not null"`
// 	NumberOfPages uint   `gorm:"type:int;not null"`
// 	Status        bool   `gorm:"not null"`
// 	CreatedAt     time.Time
// 	UpdatedAt     time.Time
// 	DeletedAt     gorm.DeletedAt `gorm:"index"`
// }
type Books struct {
	ID            uint   `gorm:"primary_key"`
	BookId        string `gorm:"type:varchar(100);not null;unique"`
	ISBN          string `gorm:"type:varchar(13);not null;unique"`
	Publisher     string `gorm:"type:varchar(100);not null"`
	PublishDate   string `gorm:"type:varchar(100);not null"`
	Title         string `gorm:"type:varchar(100);not null"`
	Authors       string `gorm:"type:varchar(256);not null"`
	Description   string `gorm:"type:text;not null"`
	Language      string `gorm:"type:varchar(100);not null"`
	Picture       string `gorm:"type:text;not null"`
	MinDeposit    uint   `gorm:"not null"`
	NumberOfPages uint   `gorm:"type:int;not null"`
	Status        bool   `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (b *Books) BeforeCreate() (err error) {
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	return
}

func (b *Books) ToDomain() books.Domain {
	return books.Domain{
		Id:          b.ID,
		BookId:      b.BookId,
		ISBN:        b.ISBN,
		Publisher:   b.Publisher,
		PublishDate: b.PublishDate,
		Title:       b.Title,
		Authors:     b.Authors,
		// Category:    b.Category,
		Description:   b.Description,
		Language:      b.Language,
		Picture:       b.Picture,
		NumberOfPages: b.NumberOfPages,
		MinDeposit:    b.MinDeposit,
		Status:        b.Status,
		CreatedAt:     b.CreatedAt,
		UpdatedAt:     b.UpdatedAt,
	}
}

func FromDomain(domain books.Domain) Books {
	return Books{
		ID:          domain.Id,
		BookId:      domain.BookId,
		Publisher:   domain.Publisher,
		PublishDate: domain.PublishDate,
		ISBN:        domain.ISBN,
		Title:       domain.Title,
		Authors:     domain.Authors,
		// Category:    domain.Category,
		Description:   domain.Description,
		Language:      domain.Language,
		Picture:       domain.Picture,
		MinDeposit:    domain.MinDeposit,
		NumberOfPages: domain.NumberOfPages,
		Status:        domain.Status,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}

func ToListDomain(data []Books) []books.Domain {
	var listDomain []books.Domain
	for _, d := range data {
		listDomain = append(listDomain, d.ToDomain())
	}
	return listDomain
}

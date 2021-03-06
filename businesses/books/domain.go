package books

import (
	context "context"
	"time"
)

// type Domain struct {
// 	Id          uint
// 	BookId      string
// 	WorkId      string
// 	ISBN        string
// 	Publisher   []string
// 	PublishDate string
// 	Title       string
// 	// Category    string
// 	Description   string
// 	NumberOfPages uint
// 	MinDeposit    uint
// 	Status        bool
// 	CreatedAt     time.Time
// 	UpdatedAt     time.Time
// 	DeletedAt     time.Time
// }
type Domain struct {
	Id            uint
	BookId        string `validate:"required"`
	ISBN          string `validate:"required"`
	Publisher     string `validate:"required"`
	PublishDate   string `validate:"required"`
	Title         string `validate:"required"`
	Authors       string `validate:"required"`
	Description   string `validate:"required"`
	Language      string `validate:"required"`
	Picture       string `validate:"required"`
	NumberOfPages uint   `validate:"required"`
	MinDeposit    uint   `validate:"required"`
	Status        bool   `validate:"required"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type DomainRepository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id string) (Domain, error)
	GetByISBN(ctx context.Context, isbn string) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	UpdateStatus(ctx context.Context, bookId string, status bool) (Domain, error)
	// Update(ctx context.Context, domain Domain) (Domain, error)
	// Delete(ctx context.Context, id uint) (Domain, error)
}

type DomainService interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id string) (Domain, error)
	GetByISBN(ctx context.Context, isbn string) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	UpdateStatus(ctx context.Context, bookId string, status bool) (Domain, error)
	// Update(ctx context.Context, domain Domain) (Domain, error)
	// Delete(ctx context.Context, id uint) (Domain, error)
}

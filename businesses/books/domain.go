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
	BookId        string
	ISBN          string
	Publisher     string
	PublishDate   string
	Title         string
	Description   string
	Language      string
	Picture       string
	NumberOfPages uint
	MinDeposit    uint
	Status        bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type DomainRepository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	// Update(ctx context.Context, domain Domain) (Domain, error)
	// Delete(ctx context.Context, id uint) (Domain, error)
}

type DomainService interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	// Update(ctx context.Context, domain Domain) (Domain, error)
	// Delete(ctx context.Context, id uint) (Domain, error)
}

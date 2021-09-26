package responses

import (
	"Pinjem/businesses/books"
	"time"
)

type BookResponse struct {
	ID            uint
	BookId        string    `json:"bookId"`
	WorkId        string    `json:"workId"`
	ISBN          string    `json:"isbn"`
	Publisher     []string  `json:"publisher"`
	PublishDate   string    `json:"publishDate"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	NumberOfPages uint      `json:"numberOfPages"`
	MinDeposit    uint      `json:"minDeposit"`
	Status        bool      `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func FromDomain(domain books.Domain) BookResponse {
	return BookResponse{
		ID:            domain.Id,
		BookId:        domain.BookId,
		WorkId:        domain.WorkId,
		ISBN:          domain.ISBN,
		Publisher:     domain.Publisher,
		Title:         domain.Title,
		Description:   domain.Description,
		NumberOfPages: domain.NumberOfPages,
		MinDeposit:    domain.MinDeposit,
		Status:        domain.Status,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}

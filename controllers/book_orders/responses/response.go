package responses

import (
	bookOrders "Pinjem/businesses/book_orders"
	"time"
)

type BookOrderResponse struct {
	ID            uint      `json:"id"`
	OrderId       uint      `json:"order_id"`
	BookId        string    `json:"book_id"`
	DepositAmount uint      `json:"deposit_amount"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func FromDomain(domain bookOrders.Domain) BookOrderResponse {
	return BookOrderResponse{
		ID:            domain.Id,
		OrderId:       domain.OrderId,
		BookId:        domain.BookId,
		DepositAmount: domain.DepositAmount,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}

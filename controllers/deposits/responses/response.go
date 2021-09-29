package responses

import (
	"Pinjem/businesses/deposits"
	"time"
)

type DepositResponse struct {
	ID        uint      `json:"id"`
	UserId    uint      `json:"user_id"`
	Amount    uint      `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain deposits.Domain) DepositResponse {
	return DepositResponse{
		ID:        domain.Id,
		UserId:    domain.UserId,
		Amount:    domain.Amount,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

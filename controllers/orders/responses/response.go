package responses

import (
	"Pinjem/businesses/orders"
	"time"
)

type OrderResponse struct {
	ID            uint      `json:"id"`
	UserId        uint      `json:"user_id"`
	OrderDate     time.Time `json:"order_date"`
	ExpDate       time.Time `json:"exp_date"`
	BookId        []string  `json:"book_id"`
	DestProvinsi  string    `json:"dest_provinsi"`
	DestKota      string    `json:"dest_kota"`
	DestKecamatan string    `json:"dest_kecamatan"`
	DestDesa      string    `json:"dest_desa"`
	DestAddress   string    `json:"dest_address"`
	Status        bool      `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func FromDomain(domain orders.Domain) OrderResponse {
	return OrderResponse{
		ID:        domain.Id,
		UserId:    domain.UserId,
		OrderDate: domain.OrderDate,
		ExpDate:   domain.ExpDate,
		Status:    domain.Status,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

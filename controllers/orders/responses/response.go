package responses

import (
	"Pinjem/businesses/orders"
	"Pinjem/businesses/shipping_details"
	"time"
)

type OrderResponse struct {
	ID             uint      `json:"id"`
	UserId         uint      `json:"user_id"`
	OrderDate      time.Time `json:"order_date"`
	ExpDate        time.Time `json:"exp_date"`
	BookId         []string  `json:"book_id"`
	DestProvinsi   string    `json:"dest_provinsi"`
	DestKota       string    `json:"dest_kota"`
	DestKecamatan  string    `json:"dest_kecamatan"`
	DestDesa       string    `json:"dest_desa"`
	DestAddress    string    `json:"dest_address"`
	DestPostalCode string    `json:"dest_postal_code"`
	ShippingCost   uint      `json:"shipping_cost"`
	Status         bool      `json:"status"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func FromDomain(orderDomain orders.Domain, shipppingDetailDomain shipping_details.Domain, bookIds []string) OrderResponse {
	return OrderResponse{
		ID:             orderDomain.Id,
		UserId:         orderDomain.UserId,
		OrderDate:      orderDomain.OrderDate,
		ExpDate:        orderDomain.ExpDate,
		BookId:         bookIds,
		DestProvinsi:   shipppingDetailDomain.DestProvinsi,
		DestKota:       shipppingDetailDomain.DestKota,
		DestKecamatan:  shipppingDetailDomain.DestKecamatan,
		DestDesa:       shipppingDetailDomain.DestDesa,
		DestAddress:    shipppingDetailDomain.DestAddress,
		DestPostalCode: shipppingDetailDomain.DestPostalCode,
		ShippingCost:   shipppingDetailDomain.ShippingCost,
		Status:         orderDomain.Status,
		CreatedAt:      orderDomain.CreatedAt,
		UpdatedAt:      orderDomain.UpdatedAt,
	}
}

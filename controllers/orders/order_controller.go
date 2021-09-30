package orders

import (
	"Pinjem/businesses/orders"
	"Pinjem/controllers"
	"Pinjem/controllers/orders/requests"
	"Pinjem/controllers/orders/responses"
	"Pinjem/helpers"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	Usecase orders.Usecase
}

type Header struct {
	Cookie string `json:"cookie"`
}

func NewOrderController(u orders.Usecase) *OrderController {
	return &OrderController{
		Usecase: u,
	}
}

func (b *OrderController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	orders, err := b.Usecase.GetAll(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := make([]responses.OrderResponse, len(orders))
	for i, order := range orders {
		response[i] = responses.OrderResponse{
			ID:        order.Id,
			UserId:    order.UserId,
			OrderDate: order.OrderDate,
			ExpDate:   order.ExpDate,
			Status:    order.Status,
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
		}
	}
	return controllers.SuccessResponse(c, response)
}

func (u *OrderController) GetById(c echo.Context) error {
	ctx := c.Request().Context()

	orderId := c.Param("orderId")
	// orderIdParam := c.Param("orderId")
	// orderIdInt, _ := (strconv.Atoi(orderIdParam))
	// orderId := uint(orderIdInt)
	user, err := u.Usecase.GetById(ctx, orderId)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := responses.OrderResponse{
		ID:        user.Id,
		UserId:    user.UserId,
		OrderDate: user.OrderDate,
		ExpDate:   user.ExpDate,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return controllers.SuccessResponse(c, response)
}

func (b *OrderController) Create(c echo.Context) error {
	ctx := c.Request().Context()

	createdOrder := requests.CreateOrder{}
	c.Bind(&createdOrder)

	userId, err := helpers.ExtractJWTPayloadUserId(c)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	id := uint(userId)

	orderDomain := orders.Domain{
		UserId: id,
		Status: true,
	}
	log.Println(orderDomain)

	order, err := b.Usecase.Create(ctx, orderDomain)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	OrderResponse := responses.OrderResponse{
		ID:        order.Id,
		UserId:    order.UserId,
		OrderDate: order.OrderDate,
		ExpDate:   order.ExpDate,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}

	return controllers.SuccessResponse(c, OrderResponse)
}

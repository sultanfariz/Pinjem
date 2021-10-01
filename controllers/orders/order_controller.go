package orders

import (
	bookOrders "Pinjem/businesses/book_orders"
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
	Usecase          orders.Usecase
	BookOrderUsecase bookOrders.Usecase
}

func NewOrderController(u orders.Usecase, b bookOrders.Usecase) *OrderController {
	return &OrderController{
		Usecase:          u,
		BookOrderUsecase: b,
	}
}

func (o *OrderController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	orders, err := o.Usecase.GetAll(ctx)
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

func (o *OrderController) GetById(c echo.Context) error {
	ctx := c.Request().Context()

	orderId := c.Param("orderId")
	// orderIdParam := c.Param("orderId")
	// orderIdInt, _ := (strconv.Atoi(orderIdParam))
	// orderId := uint(orderIdInt)
	user, err := o.Usecase.GetById(ctx, orderId)
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

func (o *OrderController) Create(c echo.Context) error {
	// ctx := c.Request().Context()

	log.Println("---------------------------------")
	books := c.FormValue("books")
	log.Println(books)
	createdOrder := requests.CreateOrder{}
	c.Bind(&createdOrder)
	log.Println(createdOrder)
	log.Println(createdOrder.Books)

	// get user id from token
	userId, err := helpers.ExtractJWTPayloadUserId(c)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	id := uint(userId)

	orderDomain := orders.Domain{
		UserId: id,
		Status: true,
	}
	log.Println(createdOrder)
	log.Println(orderDomain)

	// input order to db
	// order, err := o.Usecase.Create(ctx, orderDomain)
	// if err != nil {
	// 	return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	// }

	// for _, bookId := range createdOrder.books {
	// 	bookOrderDomain := bookOrders.Domain{
	// 		OrderId: order.Id,
	// 		BookId:  bookId,
	// 	}
	// 	log.Println(bookOrderDomain)
	// }

	OrderResponse := responses.OrderResponse{
		// ID:        order.Id,
		// UserId:    order.UserId,
		// OrderDate: order.OrderDate,
		// ExpDate:   order.ExpDate,
		// Status:    order.Status,
		// CreatedAt: order.CreatedAt,
		// UpdatedAt: order.UpdatedAt,
	}

	// return controllers.SuccessResponse(c, OrderResponse)
	return controllers.SuccessResponse(c, OrderResponse)
}

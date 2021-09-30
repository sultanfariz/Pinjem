package book_orders

import (
	bookOrders "Pinjem/businesses/book_orders"
	"Pinjem/controllers"
	"Pinjem/controllers/book_orders/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookOrderController struct {
	Usecase bookOrders.Usecase
}

func NewBookOrderController(u bookOrders.Usecase) *BookOrderController {
	return &BookOrderController{
		Usecase: u,
	}
}

func (b *BookOrderController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	bookOrders, err := b.Usecase.GetAll(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := make([]responses.BookOrderResponse, len(bookOrders))
	for i, bookOrder := range bookOrders {
		response[i] = responses.FromDomain(bookOrder)
		// response[i] = responses.BookOrderResponse{
		// 	ID:        order.Id,
		// 	UserId:    order.UserId,
		// 	OrderDate: order.OrderDate,
		// 	ExpDate:   order.ExpDate,
		// 	Status:    order.Status,
		// 	CreatedAt: order.CreatedAt,
		// 	UpdatedAt: order.UpdatedAt,
		// }
	}
	return controllers.SuccessResponse(c, response)
}

func (u *BookOrderController) GetById(c echo.Context) error {
	ctx := c.Request().Context()

	bookOrderId := c.Param("bookOrderId")
	// bookOrderIdParam := c.Param("bookOrderId")
	// bookOrderIdInt, _ := (strconv.Atoi(bookOrderIdParam))
	// bookOrderId := uint(bookOrderIdInt)
	bookOrder, err := u.Usecase.GetById(ctx, bookOrderId)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	response := responses.FromDomain(bookOrder)
	// response := responses.BookOrderResponse{
	// 	ID:        bookOrder.Id,
	// 	OrderId:   bookOrder.OrderId,
	// 	BookId:    bookOrder.BookId,
	// 	CreatedAt: bookOrder.CreatedAt,
	// 	UpdatedAt: bookOrder.UpdatedAt,
	// }

	return controllers.SuccessResponse(c, response)
}

// func (b *BookOrderController) Create(c echo.Context) error {
// 	ctx := c.Request().Context()

// 	createdOrder := requests.CreateOrder{}
// 	c.Bind(&createdOrder)

// 	userId, err := helpers.ExtractJWTPayloadUserId(c)
// 	if err != nil {
// 		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	id := uint(userId)

// 	orderDomain := bookOrders.Domain{
// 		UserId: id,
// 		Status: true,
// 	}
// 	log.Println(orderDomain)

// 	order, err := b.Usecase.Create(ctx, orderDomain)
// 	if err != nil {
// 		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	bookOrderResponse := responses.BookOrderResponse{
// 		ID:        order.Id,
// 		UserId:    order.UserId,
// 		OrderDate: order.OrderDate,
// 		ExpDate:   order.ExpDate,
// 		Status:    order.Status,
// 		CreatedAt: order.CreatedAt,
// 		UpdatedAt: order.UpdatedAt,
// 	}

// 	return controllers.SuccessResponse(c, BookOrderResponse)
// }

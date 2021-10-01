package orders

import (
	bookOrders "Pinjem/businesses/book_orders"
	"Pinjem/businesses/books"
	"Pinjem/businesses/orders"
	"Pinjem/controllers"
	"Pinjem/controllers/orders/requests"
	"Pinjem/controllers/orders/responses"
	"Pinjem/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	Usecase          orders.Usecase
	BookUsecase      books.Usecase
	BookOrderUsecase bookOrders.Usecase
}

func NewOrderController(u orders.Usecase, bo bookOrders.Usecase, b books.Usecase) *OrderController {
	return &OrderController{
		Usecase:          u,
		BookOrderUsecase: bo,
		BookUsecase:      b,
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

	orderIdParam := c.Param("orderId")
	orderIdInt, _ := (strconv.Atoi(orderIdParam))
	orderId := uint(orderIdInt)
	user, err := o.Usecase.GetById(ctx, orderId)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	books, err := o.BookOrderUsecase.GetByOrderId(ctx, orderId)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	var bookIds []string
	for _, book := range books {
		bookIds = append(bookIds, book.BookId)
	}

	response := responses.OrderResponse{
		ID:        user.Id,
		UserId:    user.UserId,
		OrderDate: user.OrderDate,
		ExpDate:   user.ExpDate,
		BookId:    bookIds,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return controllers.SuccessResponse(c, response)
}

func (o *OrderController) Create(c echo.Context) error {
	ctx := c.Request().Context()

	createdOrder := requests.CreateOrder{}
	c.Bind(&createdOrder)

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

	// input order to db
	order, err := o.Usecase.Create(ctx, orderDomain)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	// input book order to db
	for _, bookId := range createdOrder.Books {
		book, err := o.BookUsecase.GetById(ctx, bookId)
		if book.Id == 0 {
			return controllers.ErrorResponse(c, http.StatusBadRequest, err)
		}
		if err != nil {
			return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
		}
		bookOrderDomain := bookOrders.Domain{
			OrderId:       order.Id,
			BookId:        bookId,
			DepositAmount: book.MinDeposit,
		}
		bookOrder, err := o.BookOrderUsecase.Create(ctx, bookOrderDomain)
		if bookOrder.Id == 0 {
			return controllers.ErrorResponse(c, http.StatusBadRequest, err)
		}
		if err != nil {
			return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
		}
	}

	OrderResponse := responses.OrderResponse{
		ID:        order.Id,
		UserId:    order.UserId,
		OrderDate: order.OrderDate,
		ExpDate:   order.ExpDate,
		BookId:    createdOrder.Books,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}

	return controllers.SuccessResponse(c, OrderResponse)
}

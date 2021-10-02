package orders

import (
	bookOrders "Pinjem/businesses/book_orders"
	"Pinjem/businesses/books"
	"Pinjem/businesses/deposits"
	"Pinjem/businesses/orders"
	"Pinjem/controllers"
	"Pinjem/controllers/orders/requests"
	"Pinjem/controllers/orders/responses"
	"Pinjem/exceptions"
	"Pinjem/helpers"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	Usecase          orders.Usecase
	BookUsecase      books.Usecase
	BookOrderUsecase bookOrders.Usecase
	DepositUsecase   deposits.Usecase
}

func NewOrderController(u orders.Usecase, bo bookOrders.Usecase, b books.Usecase, d deposits.Usecase) *OrderController {
	return &OrderController{
		Usecase:          u,
		BookOrderUsecase: bo,
		BookUsecase:      b,
		DepositUsecase:   d,
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
		books, err := o.BookOrderUsecase.GetByOrderId(ctx, order.Id)
		if err != nil {
			return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
		}

		var bookIds []string
		for _, book := range books {
			bookIds = append(bookIds, book.BookId)
		}

		response[i] = responses.OrderResponse{
			ID:        order.Id,
			UserId:    order.UserId,
			OrderDate: order.OrderDate,
			ExpDate:   order.ExpDate,
			BookId:    bookIds,
			Status:    order.Status,
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
		}
	}
	return controllers.SuccessResponse(c, response)
}

func (o *OrderController) GetMyOrders(c echo.Context) error {
	ctx := c.Request().Context()

	userId, err := helpers.ExtractJWTPayloadUserId(c)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	id := uint(userId)

	orders, err := o.Usecase.GetOrdersByUserId(ctx, id)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := make([]responses.OrderResponse, len(orders))
	for i, order := range orders {
		books, err := o.BookOrderUsecase.GetByOrderId(ctx, order.Id)
		if err != nil {
			return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
		}

		var bookIds []string
		for _, book := range books {
			bookIds = append(bookIds, book.BookId)
		}

		response[i] = responses.OrderResponse{
			ID:        order.Id,
			UserId:    order.UserId,
			OrderDate: order.OrderDate,
			ExpDate:   order.ExpDate,
			BookId:    bookIds,
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
	var totalDeposit uint
	var bookOrderDomainArr []bookOrders.Domain
	for _, bookId := range createdOrder.Books {
		// check if book available and get book price
		book, err := o.BookUsecase.GetById(ctx, bookId)
		if book.Id == 0 || !book.Status {
			log.Println(book.Id, book.Status)
			_ = o.Usecase.Delete(ctx, order.Id)
			return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBookNotFound)
		}
		if err != nil {
			err = o.Usecase.Delete(ctx, order.Id)
			return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
		}

		// insert to book order db
		bookOrderDomain := bookOrders.Domain{
			OrderId:       order.Id,
			BookId:        bookId,
			DepositAmount: book.MinDeposit,
		}
		bookOrderDomainArr = append(bookOrderDomainArr, bookOrderDomain)
		totalDeposit += bookOrderDomain.DepositAmount
	}

	// check if total deposit amount is enough
	deposit, err := o.DepositUsecase.GetByUserId(ctx, id)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	if deposit.Amount < totalDeposit {
		_ = o.Usecase.Delete(ctx, order.Id)
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrInsufficientBalance)
	}

	for _, bookOrderDomain := range bookOrderDomainArr {
		bookOrder, err := o.BookOrderUsecase.Create(ctx, bookOrderDomain)
		if bookOrder.Id == 0 {
			return controllers.ErrorResponse(c, http.StatusBadRequest, err)
		}
		if err != nil {
			return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
		}

		// update book status
		_, err = o.BookUsecase.UpdateStatus(ctx, bookOrderDomain.BookId, false)
		if err != nil {
			return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
		}
	}

	// update deposit amount
	_, err = o.DepositUsecase.Update(ctx, id, deposit.Amount-totalDeposit)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
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

func (o *OrderController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	orderIdParam := c.Param("orderId")
	orderIdInt, _ := (strconv.Atoi(orderIdParam))
	orderId := uint(orderIdInt)

	// get order by id
	order, err := o.Usecase.GetById(ctx, orderId)
	if order.Id == 0 {
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrOrderNotFound)
	}
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	// get book order by order id
	bookOrders, err := o.BookOrderUsecase.GetByOrderId(ctx, orderId)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	// get book id from book order
	var bookIds []string
	for _, bookOrder := range bookOrders {
		bookIds = append(bookIds, bookOrder.BookId)
	}

	// update book status
	for _, bookId := range bookIds {
		_, err := o.BookUsecase.UpdateStatus(ctx, bookId, true)
		if err != nil {
			return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
		}
	}

	// delete book order
	err = o.BookOrderUsecase.DeleteByOrderId(ctx, orderId)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	// delete order
	err = o.Usecase.Delete(ctx, orderId)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.SuccessResponse(c, nil)
}

package orders

import (
	bookOrders "Pinjem/businesses/book_orders"
	"Pinjem/businesses/books"
	"Pinjem/businesses/deposits"
	"Pinjem/businesses/orders"
	shippingDetails "Pinjem/businesses/shipping_details"
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
	Usecase               orders.Usecase
	BookUsecase           books.Usecase
	BookOrderUsecase      bookOrders.Usecase
	DepositUsecase        deposits.Usecase
	ShippingDetailUsecase shippingDetails.Usecase
}

func NewOrderController(u orders.Usecase, bo bookOrders.Usecase, b books.Usecase, d deposits.Usecase, s shippingDetails.Usecase) *OrderController {
	return &OrderController{
		Usecase:               u,
		BookOrderUsecase:      bo,
		BookUsecase:           b,
		DepositUsecase:        d,
		ShippingDetailUsecase: s,
	}
}

func (o *OrderController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	orders, err := o.Usecase.GetAll(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	response := make([]responses.OrderResponse, len(orders))
	for i, order := range orders {
		books, err := o.BookOrderUsecase.GetByOrderId(ctx, order.Id)
		if err != nil {
			return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
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
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	id := uint(userId)

	orders, err := o.Usecase.GetOrdersByUserId(ctx, id)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	response := make([]responses.OrderResponse, len(orders))
	for i, order := range orders {
		// get shipping details data
		shippingDetails, err := o.ShippingDetailUsecase.GetByOrderId(ctx, order.Id)
		if err != nil {
			return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		}

		// get books data
		books, err := o.BookOrderUsecase.GetByOrderId(ctx, order.Id)
		if err != nil {
			return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		}

		var bookIds []string
		for _, book := range books {
			bookIds = append(bookIds, book.BookId)
		}
		log.Println(shippingDetails)
		response[i] = responses.FromDomain(order, shippingDetails, bookIds)
	}
	return controllers.SuccessResponse(c, response)
}

func (o *OrderController) GetById(c echo.Context) error {
	ctx := c.Request().Context()

	orderIdParam := c.Param("orderId")
	orderIdInt, _ := (strconv.Atoi(orderIdParam))
	orderId := uint(orderIdInt)
	order, err := o.Usecase.GetById(ctx, orderId)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	// get books data
	books, err := o.BookOrderUsecase.GetByOrderId(ctx, orderId)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	var bookIds []string
	for _, book := range books {
		bookIds = append(bookIds, book.BookId)
	}

	// get shipping detail data
	shippingDetail, err := o.ShippingDetailUsecase.GetByOrderId(ctx, orderId)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	response := responses.FromDomain(order, shippingDetail, bookIds)

	return controllers.SuccessResponse(c, response)
}

func (o *OrderController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	createdOrder := requests.CreateOrder{}
	c.Bind(&createdOrder)

	// get user id from token
	userId, err := helpers.ExtractJWTPayloadUserId(c)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	id := uint(userId)

	// input order to db
	orderDomain := orders.Domain{
		UserId: id,
		Status: true,
	}
	order, err := o.Usecase.Create(ctx, orderDomain)
	if order.Id == 0 {
		return controllers.ErrorResponse(c, http.StatusBadRequest, err)
	}
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	// store shipping detail to db
	shippingDetailDomain := shippingDetails.Domain{
		OrderId:        order.Id,
		DestProvinsi:   createdOrder.DestProvinsi,
		DestKota:       createdOrder.DestKota,
		DestKecamatan:  createdOrder.DestKecamatan,
		DestDesa:       createdOrder.DestDesa,
		DestAddress:    createdOrder.DestAddress,
		DestPostalCode: createdOrder.DestPostalCode,
		ShippingCost:   createdOrder.ShippingCost,
	}
	shippingDetail, err := o.ShippingDetailUsecase.Create(ctx, shippingDetailDomain)
	if shippingDetail.Id == 0 {
		_ = o.Usecase.Delete(ctx, order.Id)
		return controllers.ErrorResponse(c, http.StatusBadRequest, err)
	}
	if err != nil {
		_ = o.Usecase.Delete(ctx, order.Id)
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	// input book order to db
	var totalDeposit uint
	var bookOrderDomainArr []bookOrders.Domain
	for _, bookId := range createdOrder.Books {
		// check if book available and get book price
		book, err := o.BookUsecase.GetById(ctx, bookId)
		if book.Id == 0 {
			_ = o.Usecase.Delete(ctx, order.Id)
			_ = o.ShippingDetailUsecase.Delete(ctx, shippingDetail.Id)
			return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBookNotFound)
		}
		if !book.Status {
			_ = o.Usecase.Delete(ctx, order.Id)
			_ = o.ShippingDetailUsecase.Delete(ctx, shippingDetail.Id)
			return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBookNotAvailable)
		}
		if err != nil {
			_ = o.Usecase.Delete(ctx, order.Id)
			_ = o.ShippingDetailUsecase.Delete(ctx, shippingDetail.Id)
			return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
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
		_ = o.Usecase.Delete(ctx, order.Id)
		_ = o.ShippingDetailUsecase.Delete(ctx, shippingDetail.Id)
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	cost := shippingDetailDomain.ShippingCost + totalDeposit
	if deposit.Amount < cost {
		_ = o.Usecase.Delete(ctx, order.Id)
		_ = o.ShippingDetailUsecase.Delete(ctx, shippingDetail.Id)
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrInsufficientBalance)
	}

	for _, bookOrderDomain := range bookOrderDomainArr {
		bookOrder, err := o.BookOrderUsecase.Create(ctx, bookOrderDomain)
		if bookOrder.Id == 0 {
			return controllers.ErrorResponse(c, http.StatusBadRequest, err)
		}
		if err != nil {
			return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		}

		// update book status
		_, err = o.BookUsecase.UpdateStatus(ctx, bookOrderDomain.BookId, false)
		if err != nil {
			return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		}
	}

	// update deposit amount
	_, err = o.DepositUsecase.Update(ctx, id, deposit.Amount-totalDeposit, totalDeposit)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	OrderResponse := responses.FromDomain(order, shippingDetail, createdOrder.Books)

	return controllers.SuccessResponse(c, OrderResponse)
}

func (o *OrderController) UpdateStatus(c echo.Context) error {
	ctx := c.Request().Context()
	orderIdParam := c.Param("orderId")
	orderIdInt, _ := (strconv.Atoi(orderIdParam))
	orderId := uint(orderIdInt)
	var updateOrderStatus requests.UpdateOrderStatus
	c.Bind(&updateOrderStatus)

	order, err := o.Usecase.UpdateStatus(ctx, orderId, updateOrderStatus.Status)
	if order.Id == 0 {
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrOrderNotFound)
	}
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	// get books data
	books, err := o.BookOrderUsecase.GetByOrderId(ctx, orderId)
	if err != nil {
		_, _ = o.Usecase.UpdateStatus(ctx, orderId, !updateOrderStatus.Status)
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	var bookIds []string
	for _, book := range books {
		bookIds = append(bookIds, book.BookId)
	}

	var totalRefund uint = 0
	// update book status
	for _, bookId := range bookIds {
		_, err := o.BookUsecase.UpdateStatus(ctx, bookId, true)
		if err != nil {
			return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		}
		book, err := o.BookUsecase.GetById(ctx, bookId)
		if err != nil {
			return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		}
		totalRefund += book.MinDeposit
	}

	// get shipping detail data
	shippingDetail, err := o.ShippingDetailUsecase.GetByOrderId(ctx, orderId)
	if shippingDetail.Id == 0 {
		_, _ = o.Usecase.UpdateStatus(ctx, orderId, !updateOrderStatus.Status)
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrShippingDetailNotFound)
	}
	if err != nil {
		_, _ = o.Usecase.UpdateStatus(ctx, orderId, !updateOrderStatus.Status)
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	// update deposit amount
	_, err = o.DepositUsecase.TopUp(ctx, order.UserId, totalRefund)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	OrderResponse := responses.OrderResponse{
		ID:        order.Id,
		UserId:    order.UserId,
		OrderDate: order.OrderDate,
		ExpDate:   order.ExpDate,
		BookId:    bookIds,
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
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	// get book order by order id
	bookOrders, err := o.BookOrderUsecase.GetByOrderId(ctx, orderId)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	// get book id from book order
	var bookIds []string
	for _, bookOrder := range bookOrders {
		bookIds = append(bookIds, bookOrder.BookId)
	}

	var totalRefund uint = 0
	// update book status
	for _, bookId := range bookIds {
		_, err := o.BookUsecase.UpdateStatus(ctx, bookId, true)
		if err != nil {
			return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		}
		book, err := o.BookUsecase.GetById(ctx, bookId)
		if err != nil {
			return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		}
		totalRefund += book.MinDeposit
	}

	// delete book order
	err = o.BookOrderUsecase.DeleteByOrderId(ctx, orderId)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	// delete shipping detail
	shippingDetail, err := o.ShippingDetailUsecase.GetByOrderId(ctx, orderId)
	if shippingDetail.Id == 0 {
		return controllers.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrShippingDetailNotFound)
	}
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	err = o.ShippingDetailUsecase.DeleteByOrderId(ctx, orderId)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	// delete order
	err = o.Usecase.Delete(ctx, orderId)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	// update deposit amount
	_, err = o.DepositUsecase.TopUp(ctx, order.UserId, (shippingDetail.ShippingCost + totalRefund))
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	return controllers.SuccessResponse(c, nil)
}

package book_orders_test

import (
	bookOrders "Pinjem/businesses/book_orders"
	"Pinjem/businesses/book_orders/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var bookOrderRepository mocks.DomainRepository

var bookOrderService bookOrders.DomainService
var bookOrderDomain bookOrders.Domain

func setup() {
	bookOrderService = bookOrders.NewUsecase(&bookOrderRepository, time.Minute*15)
	bookOrderDomain = bookOrders.Domain{
		Id:            1,
		BookId:        "DlQbmJc5WlQC",
		OrderId:       1,
		DepositAmount: 100000,
	}
}

func TestGetAllBookOrders(t *testing.T) {
	setup()
	bookOrderRepository.On("GetAll", mock.Anything).Return([]bookOrders.Domain{bookOrderDomain}, nil)
	t.Run("Test case 1 | Valid Get All Book Orders", func(t *testing.T) {
		bookOrders, err := bookOrderService.GetAll(context.Background())
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if len(bookOrders) == 0 {
			t.Errorf("Error: %s", "No book orders found")
		}

		assert.Nil(t, err)
		assert.Equal(t, bookOrderDomain, bookOrders[0])
	})
}

func TestGetBookOrderByOrderId(t *testing.T) {
	setup()
	bookOrderRepository.On("GetByOrderId", mock.Anything, mock.AnythingOfType("uint")).Return([]bookOrders.Domain{bookOrderDomain}, nil)
	t.Run("Test case 1 | Valid Get Book Order By Order Id", func(t *testing.T) {
		bookOrders, err := bookOrderService.GetByOrderId(context.Background(), 1)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if bookOrders == nil {
			t.Errorf("Error: %s", "No book order found")
		}

		assert.Nil(t, err)
		assert.Equal(t, bookOrderDomain, bookOrders[0])
	})
}

func TestCreateNewBookOrder(t *testing.T) {
	setup()
	bookOrderRepository.On("Create", mock.Anything, mock.AnythingOfType("Domain")).Return(bookOrderDomain, nil)
	t.Run("Test case 1 | Valid Create New Book Order", func(t *testing.T) {
		bookOrder, err := bookOrderService.Create(context.Background(), bookOrderDomain)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if bookOrder.Id == 0 {
			t.Errorf("Error: %s", "No book order found")
		}

		assert.Nil(t, err)
		assert.Equal(t, bookOrderDomain, bookOrder)
	})
}

func TestDeleteBookOrderByOrderId(t *testing.T) {
	setup()
	bookOrderRepository.On("DeleteByOrderId", mock.Anything, mock.AnythingOfType("uint")).Return(nil)
	t.Run("Test case 1 | Valid Delete Book Order By Order Id", func(t *testing.T) {
		err := bookOrderService.DeleteByOrderId(context.Background(), 1)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}

		assert.Nil(t, err)
	})
}

package orders_test

import (
	"Pinjem/businesses/orders"
	"Pinjem/businesses/orders/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var orderRepository mocks.DomainRepository

var orderService orders.DomainService
var orderDomain orders.Domain

func setup() {
	orderService = orders.NewUsecase(&orderRepository, time.Minute*15)
	orderDomain = orders.Domain{
		Id:        1,
		UserId:    1,
		OrderDate: time.Now(),
		ExpDate:   time.Now().AddDate(0, 0, 60),
		Status:    true,
	}
}

func TestGetAllOrders(t *testing.T) {
	setup()
	orderRepository.On("GetAll", mock.Anything).Return([]orders.Domain{orderDomain}, nil)
	t.Run("Test Case 1 | Valid Get All Orders", func(t *testing.T) {
		orders, err := orderService.GetAll(context.Background())
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if len(orders) < 1 {
			t.Errorf("Error: %s", "No orders found")
		}

		assert.Nil(t, err)
		assert.Equal(t, orderDomain, orders[0])
	})
}

func TestGetOrderById(t *testing.T) {
	setup()
	orderRepository.On("GetById", mock.Anything, mock.Anything).Return(orderDomain, nil)
	t.Run("Test Case 1 | Valid Get Order By Id", func(t *testing.T) {
		order, err := orderService.GetById(context.Background(), 1)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if order.Id != 1 {
			t.Errorf("Error: %s", "No order found")
		}

		assert.Nil(t, err)
		assert.Equal(t, orderDomain, order)
	})
}

func TestGetOrdersByUserId(t *testing.T) {
	setup()
	orderRepository.On("GetOrdersByUserId", mock.Anything, mock.AnythingOfType("uint")).Return([]orders.Domain{orderDomain}, nil)
	t.Run("Test Case 1 | Valid Get Orders By User Id", func(t *testing.T) {
		orders, err := orderService.GetOrdersByUserId(context.Background(), 1)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if len(orders) < 1 {
			t.Errorf("Error: %s", "No orders found")
		}

		assert.Nil(t, err)
		assert.Equal(t, orderDomain, orders[0])
	})
}

func TestCreateOrder(t *testing.T) {
	setup()
	orderRepository.On("Create", mock.Anything, mock.AnythingOfType("Domain")).Return(orderDomain, nil)
	t.Run("Test Case 1 | Valid Create Order", func(t *testing.T) {
		order, err := orderService.Create(context.Background(), orderDomain)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if order.Id != 1 {
			t.Errorf("Error: %s", "No order found")
		}

		assert.Nil(t, err)
		assert.Equal(t, orderDomain, order)
	})
}

// func TestUpdateOrder(t *testing.T) {
// 	setup()
// 	orderRepository.On("UpdateStatus", mock.Anything, mock.AnythingOfType("uint"), mock.AnythingOfType("bool")).Return(orderDomain, nil)
// 	t.Run("Test Case 1 | Valid Update Order Status", func(t *testing.T) {
// 		order, err := orderService.UpdateStatus(context.Background(), orderDomain.Id, false)
// 		if err != nil {
// 			t.Errorf("Error: %s", err.Error())
// 		}
// 		if order.Id == 0 {
// 			t.Errorf("Error: %s", "No order found")
// 		}

// 		assert.Nil(t, err)
// 		assert.Equal(t, false, order.Status)
// 	})
// }

func TestDeleteOrder(t *testing.T) {
	setup()
	orderRepository.On("Delete", mock.Anything, mock.AnythingOfType("uint")).Return(nil)
	t.Run("Test Case 1 | Valid Delete Order", func(t *testing.T) {
		err := orderService.Delete(context.Background(), orderDomain.Id)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}

		assert.Nil(t, err)
	})
}

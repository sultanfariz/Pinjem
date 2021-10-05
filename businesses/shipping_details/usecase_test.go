package shipping_details_test

import (
	shippingDetails "Pinjem/businesses/shipping_details"
	"Pinjem/businesses/shipping_details/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var shippingDetailRepository mocks.DomainRepository

var shippingDetailService shippingDetails.DomainService
var shippingDetailDomain shippingDetails.Domain

func setup() {
	shippingDetailService = shippingDetails.NewUsecase(&shippingDetailRepository, time.Minute*15)
	shippingDetailDomain = shippingDetails.Domain{
		Id:             1,
		OrderId:        1,
		DestProvinsi:   "Jawa Barat",
		DestKota:       "Bandung",
		DestKecamatan:  "Cibadak",
		DestDesa:       "Cibadak",
		DestAddress:    "Jl. Cibadak",
		DestPostalCode: "40132",
		ShippingCost:   9000,
	}
}

func TestGetAllShippingDetails(t *testing.T) {
	setup()
	shippingDetailRepository.On("GetAll", mock.Anything).Return([]shippingDetails.Domain{shippingDetailDomain}, nil)
	t.Run("Test Case 1 | Valid Get All Shipping Details", func(t *testing.T) {
		shippingDetails, err := shippingDetailService.GetAll(context.Background())
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if len(shippingDetails) < 1 {
			t.Errorf("Error: %v", "No Shipping Details")
		}

		assert.NoError(t, err)
		assert.Equal(t, 1, len(shippingDetails))
	})
}
func TestGetShippingDetailById(t *testing.T) {
	setup()
	shippingDetailRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(shippingDetailDomain, nil)
	t.Run("Test Case 1 | Valid Get Shipping Detail By Id", func(t *testing.T) {
		shippingDetail, err := shippingDetailService.GetById(context.Background(), 1)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if shippingDetail.Id == 0 {
			t.Errorf("Error: %v", "No Shipping Details")
		}

		assert.NoError(t, err)
		assert.Equal(t, uint(1), shippingDetail.Id)
	})
}

func TestGetShippingDetailByOrderId(t *testing.T) {
	setup()
	shippingDetailRepository.On("GetByOrderId", mock.Anything, mock.Anything).Return(shippingDetailDomain, nil)
	t.Run("Test Case 1 | Valid Get Shipping Detail By Order Id", func(t *testing.T) {
		shippingDetail, err := shippingDetailService.GetByOrderId(context.Background(), 1)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if shippingDetail.Id < 1 {
			t.Errorf("Error: %v", "No Shipping Details")
		}

		assert.NoError(t, err)
		assert.Equal(t, uint(1), shippingDetail.Id)
	})
}

func TestCreateShippingDetail(t *testing.T) {
	setup()
	shippingDetailRepository.On("Create", mock.Anything, mock.AnythingOfType("Domain")).Return(shippingDetailDomain, nil)
	t.Run("Test Case 1 | Valid Create Shipping Detail", func(t *testing.T) {
		shippingDetail, err := shippingDetailService.Create(context.Background(), shippingDetailDomain)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if shippingDetail.Id < 1 {
			t.Errorf("Error: %v", "No Shipping Details")
		}

		assert.NoError(t, err)
		assert.Equal(t, shippingDetailDomain, shippingDetail)
	})
}

func TestDeleteShippingDetail(t *testing.T) {
	setup()
	shippingDetailRepository.On("Delete", mock.Anything, mock.AnythingOfType("uint")).Return(nil)
	t.Run("Test Case 1 | Valid Delete Shipping Detail", func(t *testing.T) {
		err := shippingDetailService.Delete(context.Background(), 1)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		assert.NoError(t, err)
	})
}

func TestDeleteShippingDetailByOrderId(t *testing.T) {
	setup()
	shippingDetailRepository.On("DeleteByOrderId", mock.Anything, mock.AnythingOfType("uint")).Return(nil)
	t.Run("Test Case 1 | Valid Delete Shipping Detail By Order Id", func(t *testing.T) {
		err := shippingDetailService.DeleteByOrderId(context.Background(), 1)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		assert.NoError(t, err)
	})
}

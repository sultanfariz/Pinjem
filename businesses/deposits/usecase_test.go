package deposits_test

import (
	"Pinjem/businesses/deposits"
	"Pinjem/businesses/deposits/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var depositRepository mocks.DomainRepository

var depositService deposits.DomainService
var depositDomain deposits.Domain
var updatedDepositDomain deposits.Domain

func setup() {
	depositService = deposits.NewUsecase(&depositRepository, time.Minute*15)
	depositDomain = deposits.Domain{
		Id:         1,
		UserId:     1,
		Amount:     100000,
		UsedAmount: 0,
	}

	updatedDepositDomain = deposits.Domain{
		Id:         1,
		UserId:     1,
		Amount:     150000,
		UsedAmount: 50000,
	}
}

func TestGetAllDeposits(t *testing.T) {
	setup()
	depositRepository.On("GetAll", mock.Anything).Return([]deposits.Domain{depositDomain}, nil)
	t.Run("Test Case 1 | Get All Deposits", func(t *testing.T) {
		deposits, err := depositService.GetAll(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 1, len(deposits))
		assert.Equal(t, depositDomain, deposits[0])
	})
}

func TestGetDepositByUserId(t *testing.T) {
	setup()
	depositRepository.On("GetByUserId", mock.Anything, mock.AnythingOfType("uint")).Return(depositDomain, nil)
	t.Run("Test Case 1 | Valid Get Deposit By User Id", func(t *testing.T) {
		deposit, err := depositService.GetByUserId(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, depositDomain.UserId, deposit.UserId)
	})
}

func TestCreateNewDeposit(t *testing.T) {
	setup()
	depositRepository.On("Create", mock.Anything, mock.AnythingOfType("Domain")).Return(depositDomain, nil)
	t.Run("Test Case 1 | Valid Create New Deposit", func(t *testing.T) {
		deposit, err := depositService.Create(context.Background(), depositDomain)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if deposit.Id == 0 {
			t.Errorf("Error: %s", "Deposit Id is empty")
		}

		assert.NoError(t, err)
		assert.Equal(t, depositDomain, deposit)
	})
}

func TestUpdateDeposit(t *testing.T) {
	setup()
	depositRepository.On("Update", mock.Anything, mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(updatedDepositDomain, nil)
	t.Run("Test Case 1 | Valid Update Deposit", func(t *testing.T) {
		deposit, err := depositService.Update(context.Background(), updatedDepositDomain.UserId, updatedDepositDomain.Amount, updatedDepositDomain.UsedAmount)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if deposit.Id == 0 {
			t.Errorf("Error: %s", "Deposit Id is empty")
		}

		assert.NoError(t, err)
		assert.Equal(t, updatedDepositDomain, deposit)
	})
}

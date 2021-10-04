package users_test

import (
	"Pinjem/businesses/users"
	"Pinjem/businesses/users/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository mocks.DomainRepository

var userService users.DomainService
var userDomain users.Domain

func setup() {
	userService = users.NewUsecase(&userRepository, time.Minute*15)
	userDomain = users.Domain{
		Id:          1,
		Fullname:    "John Doe",
		Email:       "johndoe@gmail.com",
		Password:    "12345678",
		Nik:         "123456789",
		PhoneNumber: "08123456789",
		Address:     "Jl. Kebon Jeruk No. 1",
		Birthdate:   "1990-01-01",
		Provinsi:    "Jawa Barat",
		Kota:        "Bandung",
		Kecamatan:   "Kebon Jeruk",
		Desa:        "Kebon Jeruk",
		PostalCode:  "40132",
		Role:        "user",
		Status:      1,
		LinkKTP:     "http://localhost:8080/api/v1/uploads/KTP/KTP_Sultan_Fariz_20210929161427.jpg",
	}
}

func TestLogin(t *testing.T) {
	setup()
	userRepository.On("Login", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userDomain, nil).Once()
	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		user, err := userService.Login(context.Background(), userDomain.Email, userDomain.Password)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if user.Id != 1 {
			t.Errorf("Expected: %d, got: %d", 1, user.Id)
		}
		assert.Nil(t, err)
		assert.Equal(t, userDomain.Fullname, user.Fullname)
	})
	t.Run("Test Case 2 | Invalid Login with Empty Email Field", func(t *testing.T) {
		_, err := userService.Login(context.Background(), "", userDomain.Password)

		assert.NotNil(t, err)
	})
	t.Run("Test Case 3 | Invalid Login with Empty Password Field", func(t *testing.T) {
		_, err := userService.Login(context.Background(), userDomain.Email, "")

		assert.NotNil(t, err)
	})
}

func TestCreate(t *testing.T) {
	setup()
	userRepository.On("Create", mock.Anything, mock.AnythingOfType("Domain")).Return(userDomain, nil).Once()
	t.Run("Test Case 1 | Valid Register", func(t *testing.T) {
		user, err := userService.Register(context.Background(), userDomain)

		assert.Nil(t, err)
		assert.Equal(t, userDomain.Email, user.Email)
	})
	t.Run("Test Case 2 | Invalid Register with Empty Field", func(t *testing.T) {
		_, err := userService.Register(context.Background(), users.Domain{
			Fullname:    "",
			Email:       "",
			Password:    "",
			Nik:         "",
			PhoneNumber: "",
			Address:     "",
			Birthdate:   "",
			Provinsi:    "",
			Kota:        "",
			Kecamatan:   "",
			Desa:        "",
			PostalCode:  "",
			Role:        "",
			Status:      0,
			LinkKTP:     "",
		})

		assert.NotNil(t, err)
	})
}

func TestFindUserByEmail(t *testing.T) {
	setup()
	userRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(userDomain, nil).Once()
	t.Run("Test Case 1 | Valid Find User By Email", func(t *testing.T) {
		user, err := userService.FindByEmail(context.Background(), userDomain.Email)

		assert.Nil(t, err)
		assert.Equal(t, userDomain.Fullname, user.Fullname)
	})
	t.Run("Test Case 2 | Invalid Find User By Email with Empty Email Field", func(t *testing.T) {
		_, err := userService.FindByEmail(context.Background(), "")

		assert.NotNil(t, err)
	})
}

func TestGetAllUsers(t *testing.T) {
	setup()
	userRepository.On("GetAll", mock.Anything).Return([]users.Domain{userDomain}, nil).Once()
	t.Run("Test Case 1 | Get All Users", func(t *testing.T) {
		users, err := userService.GetAll(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, userDomain.Fullname, users[0].Fullname)
	})
}

func TestGetUserById(t *testing.T) {
	setup()
	userRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(userDomain, nil).Once()
	t.Run("Test Case 1 | Get User By Id", func(t *testing.T) {
		user, err := userService.GetById(context.Background(), 1)

		assert.Nil(t, err)
		assert.Equal(t, userDomain.Fullname, user.Fullname)
	})
}

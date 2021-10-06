package users

import (
	"Pinjem/businesses/deposits"
	"Pinjem/businesses/users"
	"Pinjem/controllers"
	"Pinjem/controllers/auth/responses"
	"Pinjem/exceptions"
	"Pinjem/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Usecase        users.Usecase
	DepositUsecase deposits.Usecase
}

func NewUserController(u users.Usecase, d deposits.Usecase) *UserController {
	return &UserController{
		Usecase:        u,
		DepositUsecase: d,
	}
}

func (u *UserController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	users, err := u.Usecase.GetAll(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	response := make([]responses.UserResponse, len(users))
	for i, user := range users {
		response[i] = responses.UserResponse{
			ID:          user.Id,
			Email:       user.Email,
			Fullname:    user.Fullname,
			NIK:         user.Nik,
			PhoneNumber: user.PhoneNumber,
			Birthdate:   user.Birthdate,
			Address:     user.Address,
			Provinsi:    user.Provinsi,
			Kota:        user.Kota,
			Kecamatan:   user.Kecamatan,
			Desa:        user.Desa,
			PostalCode:  user.PostalCode,
			Role:        user.Role,
			Status:      user.Status,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
		}
	}
	return controllers.SuccessResponse(c, response)
}

func (u *UserController) GetById(c echo.Context) error {
	ctx := c.Request().Context()

	idParam := c.Param("userId")
	idInt, _ := strconv.Atoi(idParam)
	id := uint(idInt)
	user, err := u.Usecase.GetById(ctx, id)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	response := responses.UserResponse{
		ID:          user.Id,
		Email:       user.Email,
		Fullname:    user.Fullname,
		NIK:         user.Nik,
		PhoneNumber: user.PhoneNumber,
		Birthdate:   user.Birthdate,
		Address:     user.Address,
		Provinsi:    user.Provinsi,
		Kota:        user.Kota,
		Kecamatan:   user.Kecamatan,
		Desa:        user.Desa,
		PostalCode:  user.PostalCode,
		Role:        user.Role,
		Status:      user.Status,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	return controllers.SuccessResponse(c, response)
}

func (u *UserController) GetMyUserProfile(c echo.Context) error {
	ctx := c.Request().Context()

	userId, err := helpers.ExtractJWTPayloadUserId(c)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	id := uint(userId)
	user, err := u.Usecase.GetById(ctx, id)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	deposit, err := u.DepositUsecase.GetByUserId(ctx, id)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	response := responses.MyProfileResponse{
		ID:            user.Id,
		Email:         user.Email,
		Fullname:      user.Fullname,
		NIK:           user.Nik,
		PhoneNumber:   user.PhoneNumber,
		Birthdate:     user.Birthdate,
		Address:       user.Address,
		Provinsi:      user.Provinsi,
		Kota:          user.Kota,
		Kecamatan:     user.Kecamatan,
		Desa:          user.Desa,
		PostalCode:    user.PostalCode,
		Role:          user.Role,
		Status:        user.Status,
		DepositAmount: deposit.Amount,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}

	return controllers.SuccessResponse(c, response)
}

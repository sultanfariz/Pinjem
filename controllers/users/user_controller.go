package users

import (
	"Pinjem/businesses/users"
	"Pinjem/controllers"
	"Pinjem/controllers/auth/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Usecase users.Usecase
}

func NewUserController(u users.Usecase) *UserController {
	return &UserController{
		Usecase: u,
	}
}

func (u *UserController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	users, err := u.Usecase.GetAll(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
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
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
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

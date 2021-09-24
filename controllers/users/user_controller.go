package users

import (
	"Pinjem/businesses/users"
	"Pinjem/controllers"
	"Pinjem/controllers/auth/responses"
	"net/http"

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

func (a *UserController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	users, err := a.Usecase.GetAll(ctx)
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

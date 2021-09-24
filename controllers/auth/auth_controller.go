package auth

import (
	"Pinjem/businesses/users"
	"Pinjem/controllers"
	"Pinjem/controllers/auth/requests"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Usecase users.Usecase
}

func NewAuthController(u users.Usecase) *AuthController {
	return &AuthController{
		Usecase: u,
	}
}

func (a *AuthController) Login(c echo.Context) error {
	userLogin := requests.LoginRequest{}
	c.Bind(&userLogin)

	ctx := c.Request().Context()
	user, err := a.Usecase.Login(ctx, userLogin.Email, userLogin.Password)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.SuccessResponse(c, user)
}

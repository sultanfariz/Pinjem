package auth

import (
	"Pinjem/businesses/users"
	"Pinjem/controllers"
	"Pinjem/controllers/auth/requests"
	"Pinjem/controllers/auth/responses"
	"Pinjem/exceptions"
	"Pinjem/helpers"
	"net/http"
	"time"

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

	// check email and password
	user, err := a.Usecase.Login(ctx, userLogin.Email, userLogin.Password)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	if user.Id == 0 {
		return controllers.ErrorResponse(c, http.StatusUnauthorized, exceptions.ErrInvalidCredentials)
	}

	// generate token and cookie
	token, err := helpers.GenerateToken(int(user.Id))
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	expirationTime := time.Now().Add(time.Hour * 24)
	helpers.SetTokenCookie("token", token, expirationTime, c)
	return controllers.SuccessResponse(c, responses.LoginResponse{Token: token})
}

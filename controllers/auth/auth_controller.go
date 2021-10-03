package auth

import (
	"Pinjem/businesses/deposits"
	"Pinjem/businesses/users"
	"Pinjem/controllers"
	"Pinjem/controllers/auth/requests"
	"Pinjem/controllers/auth/responses"
	"Pinjem/exceptions"
	"Pinjem/helpers"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Usecase        users.Usecase
	DepositUsecase deposits.Usecase
}

func NewAuthController(u users.Usecase, d deposits.Usecase) *AuthController {
	return &AuthController{
		Usecase:        u,
		DepositUsecase: d,
	}
}

func (a *AuthController) Login(c echo.Context) error {
	userLogin := requests.LoginRequest{}
	c.Bind(&userLogin)

	ctx := c.Request().Context()

	// check email and password
	user, err := a.Usecase.Login(ctx, userLogin.Email, userLogin.Password)
	if user.Id == 0 {
		return controllers.ErrorResponse(c, http.StatusUnauthorized, exceptions.ErrInvalidCredentials)
	}
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	// generate token and cookie
	token, err := helpers.GenerateToken(int(user.Id), user.Role)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	expirationTime := time.Now().Add(time.Hour * 24)
	helpers.SetTokenCookie("token", token, expirationTime, c)
	return controllers.SuccessResponse(c, responses.LoginResponse{Token: token})
}

func (a *AuthController) Register(c echo.Context) error {
	userRegister := requests.RegisterRequest{}
	c.Bind(&userRegister)

	// upload file KTP
	file, err := c.FormFile("ktp")
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	split := strings.Split(file.Filename, ".")
	extension := split[len(split)-1]
	fileName := strings.ReplaceAll(fmt.Sprintf("KTP_%s", userRegister.Fullname), " ", "_")
	filePath := "KTP"
	fileURL, fileErr := helpers.UploadFile(filePath, fileName, extension, file)
	if fileErr != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, fileErr)
	}

	userDomain := users.Domain{
		Email:       userRegister.Email,
		Password:    userRegister.Password,
		Fullname:    userRegister.Fullname,
		Nik:         userRegister.NIK,
		PhoneNumber: userRegister.PhoneNumber,
		Birthdate:   userRegister.Birthdate,
		Address:     userRegister.Address,
		Provinsi:    userRegister.Provinsi,
		Kota:        userRegister.Kota,
		Kecamatan:   userRegister.Kecamatan,
		Desa:        userRegister.Desa,
		PostalCode:  userRegister.PostalCode,
		Role:        userRegister.Role,
		Status:      userRegister.Status,
		LinkKTP:     fileURL,
	}

	ctx := c.Request().Context()

	user, err := a.Usecase.Register(ctx, userDomain)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	if user.Id == 0 {
		return controllers.ErrorResponse(c, http.StatusUnauthorized, exceptions.ErrInvalidCredentials)
	}

	// create deposit for user role
	if userRegister.Role == "user" {
		depositDomain := deposits.Domain{
			UserId: user.Id,
			Amount: 0,
		}
		_, err := a.DepositUsecase.Create(ctx, depositDomain)
		if err != nil {
			return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
		}
	}

	registerResponse := responses.RegisterResponse{
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
		LinkKTP:     user.LinkKTP,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
	return controllers.SuccessResponse(c, registerResponse)
}

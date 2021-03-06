package deposits

import (
	"Pinjem/businesses/deposits"
	"Pinjem/controllers"
	"Pinjem/controllers/deposits/responses"
	"Pinjem/exceptions"
	"Pinjem/helpers"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DepositController struct {
	Usecase deposits.Usecase
}

func NewDepositController(d deposits.Usecase) *DepositController {
	return &DepositController{
		Usecase: d,
	}
}

func (d *DepositController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	deposits, err := d.Usecase.GetAll(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	response := make([]responses.DepositResponse, len(deposits))
	for i, deposit := range deposits {
		response[i] = responses.FromDomain(deposit)
	}
	return controllers.SuccessResponse(c, response)
}

func (d *DepositController) GetByUserId(c echo.Context) error {
	ctx := c.Request().Context()

	idParam := c.Param("userId")
	log.Println(idParam)
	idInt, _ := strconv.Atoi(idParam)
	id := uint(idInt)
	deposit, err := d.Usecase.GetByUserId(ctx, id)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	response := responses.FromDomain(deposit)

	return controllers.SuccessResponse(c, response)
}

func (d *DepositController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	userId, err := helpers.ExtractJWTPayloadUserId(c)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	id := uint(userId)

	amount := c.FormValue("amount")
	amountInt, _ := strconv.Atoi(amount)
	amountUInt := uint(amountInt)

	deposit, err := d.Usecase.TopUp(ctx, id, amountUInt)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	response := responses.DepositResponse{
		ID:        deposit.Id,
		UserId:    deposit.UserId,
		Amount:    deposit.Amount,
		CreatedAt: deposit.CreatedAt,
		UpdatedAt: deposit.UpdatedAt,
	}

	return controllers.SuccessResponse(c, response)
}

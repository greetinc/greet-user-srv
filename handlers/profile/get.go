package handlers

import (
	dto "greet-user-srv/dto"

	res "github.com/greetinc/greet-util/s/response"

	"github.com/labstack/echo/v4"
)

// Get By ID
// @Summary Get By ID transaction
// @Description Get By ID transaction
// @Tags transaction
// @Accept  json
// @Produce  json
// @Security BearerUser
// @param Userorization header string true "Userorization Bearer"
// @Param id path int true "id path"
// @Success 200 {object} dto.TransactionResponse
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/transaction/{id} [get]
func (u *domainHandler) Get(c echo.Context) error {
	var req dto.UserRequest

	userId, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	req.ID = userId

	transaction, err := u.serviceUser.User(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(c)

	}

	return res.SuccessResponse(transaction).Send(c)

}

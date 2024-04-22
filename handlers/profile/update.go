package handlers

import (
	dto "greet-user-srv/dto"
	"net/http"

	res "github.com/greetinc/greet-util/s/response"

	"github.com/labstack/echo/v4"
)

// Update
// @Summary Update user
// @Description Update user
// @Tags user
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Authorization Bearer"
// @Param id path int true "id path"
// @Param request body dto.UpdateRequest true "request body"
// @Success 200 {object} dto.TransactionResponse
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/user/{id} [patch]
func (b *domainHandler) Update(c echo.Context) error {
	var req dto.UpdateUserRequest

	idUint, err := res.IsNumber(c, "id")
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	req.ID = idUint

	err = c.Bind(&req)
	if err != nil {
		return res.Response(c, http.StatusBadRequest, res.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	result, err := b.serviceUser.Update(req)
	if err != nil {
		return res.Response(c, http.StatusBadRequest, res.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return res.SuccessResponse(result).Send(c)

}

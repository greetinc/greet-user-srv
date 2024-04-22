package handlers

import (
	s "greet-user-srv/services/profile"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Get(c echo.Context) error         //cek
	UserPreview(c echo.Context) error //Preview
	Update(c echo.Context) error      //UpdateUser

}

type domainHandler struct {
	serviceUser s.UserService
}

func NewUserHandler(service s.UserService) DomainHandler {
	return &domainHandler{
		serviceUser: service,
	}
}

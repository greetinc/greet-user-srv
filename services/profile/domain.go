package user

import (
	dto "greet-user-srv/dto"
	r "greet-user-srv/repositories/profile"

	m "github.com/greetinc/greet-middlewares/middlewares"
)

type UserService interface {
	User(req dto.UserRequest) (*dto.UserResponse, error)
	UserPreview(req dto.UserPreviewRequest) (*dto.UserResponse, error)
	Update(req dto.UpdateUserRequest) (dto.UpdateUserResponse, error)
}

type userService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewUserService(Repo r.DomainRepository, jwtS m.JWTService) UserService {
	return &userService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

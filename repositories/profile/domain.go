package repositories

import (
	dto "greet-user-srv/dto"

	"github.com/greetinc/greet-auth-srv/entity"

	"sync"

	"gorm.io/gorm"
)

type DomainRepository interface {
	User(req dto.UserRequest) (*entity.User, error)
	UserPreview(req dto.UserPreviewRequest) (*entity.UserDetail, error)
	Update(req dto.UpdateUserRequest) (dto.UpdateUserResponse, error)
	GetById(req dto.UpdateUserRequest) (*dto.UpdateUserResponse, error)
}

type userRepository struct {
	DB    *gorm.DB
	mu    sync.Mutex
	users map[string]*entity.User
}

func NewUserRepository(DB *gorm.DB) DomainRepository {
	return &userRepository{
		DB: DB,
	}
}

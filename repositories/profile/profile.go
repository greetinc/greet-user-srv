package repositories

import (
	dto "greet-user-srv/dto"

	"github.com/greetinc/greet-auth-srv/entity"
)

func (u *userRepository) User(req dto.UserRequest) (*entity.User, error) {
	var existingUser entity.User
	// err := u.DB.Preload("Company").Preload("Role").Where("id = ?", req.ID).First(&existingUser).Error
	err := u.DB.Preload("File").Preload("UserDetail").Where("id = ?", req.ID).First(&existingUser).Error
	if err != nil {
		return nil, err
	}

	return &existingUser, nil
}

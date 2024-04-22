package repositories

import (
	dto "greet-user-srv/dto"

	"github.com/greetinc/greet-auth-srv/entity"
)

func (u *userRepository) UserPreview(req dto.UserPreviewRequest) (*entity.UserDetail, error) {
	var existingUser entity.UserDetail
	// err := u.DB.Preload("Company").Preload("Role").Where("id = ?", req.ID).First(&existingUser).Error
	err := u.DB.Where("user_id = ?", req.UserID).First(&existingUser).Error
	if err != nil {
		return nil, err
	}

	return &existingUser, nil
}

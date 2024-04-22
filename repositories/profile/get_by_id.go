package repositories

import (
	dto "greet-user-srv/dto"

	"github.com/greetinc/greet-auth-srv/entity"
)

func (b *userRepository) GetById(req dto.UpdateUserRequest) (*dto.UpdateUserResponse, error) {
	tr := entity.UserDetail{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.UpdateUserResponse{
		FullName:    tr.FullName,
		Description: tr.Description,
	}

	return response, nil
}

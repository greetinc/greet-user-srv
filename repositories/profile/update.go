package repositories

import (
	dto "greet-user-srv/dto"

	"github.com/greetinc/greet-auth-srv/entity"
)

func (b *userRepository) Update(req dto.UpdateUserRequest) (dto.UpdateUserResponse, error) {
	tr := dto.UpdateUserRequest{
		ID: req.ID,
	}

	request := entity.UserDetail{
		FullName:    req.FullName,
		Description: req.Description,
	}

	transaction, err := b.GetById(tr)
	if err != nil {
		return dto.UpdateUserResponse{}, err
	}

	err = b.DB.Where("ID = ?", req.ID).Updates(entity.UserDetail{
		FullName:    request.FullName,
		Description: request.Description,
	}).Error
	if err != nil {
		return dto.UpdateUserResponse{}, err
	}

	response := dto.UpdateUserResponse{
		FullName:    transaction.FullName,
		Description: transaction.Description,
	}

	return response, nil
}

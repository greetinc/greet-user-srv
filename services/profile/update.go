package user

import (
	dto "greet-user-srv/dto"
)

func (b *userService) Update(req dto.UpdateUserRequest) (dto.UpdateUserResponse, error) {
	userBody := dto.UpdateUserRequest{
		FullName:    req.FullName,
		Description: req.Description,
	}

	Transaction, err := b.Repo.Update(req)
	if err != nil {
		return Transaction, err
	}

	response := dto.UpdateUserResponse{
		FullName:    userBody.FullName,
		Description: userBody.Description,
	}

	return response, nil
}

package user

import (
	dto "greet-user-srv/dto"
	dtof "greet-user-srv/dto/upload"
)

func (u *userService) User(req dto.UserRequest) (*dto.UserResponse, error) {
	user, err := u.Repo.User(req)
	if err != nil {
		return nil, err
	}

	// Fetch file information
	var files []dtof.UploadResponse
	for _, file := range user.File {
		files = append(files, dtof.UploadResponse{
			FileName: file.FileName,
			FilePath: file.FilePath,
		})
	}

	return &dto.UserResponse{
		ID:          user.ID,
		FullName:    user.UserDetail.FullName,
		Age:         user.UserDetail.Age,
		Email:       user.Email,
		Description: user.UserDetail.Description,
		Whatsapp:    user.Whatsapp,
		Files:       files, // Menambahkan informasi file ke respons

	}, nil
}

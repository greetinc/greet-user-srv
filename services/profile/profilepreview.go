package user

import (
	dto "greet-user-srv/dto"
	dtof "greet-user-srv/dto/upload"
)

func (u *userService) UserPreview(req dto.UserPreviewRequest) (*dto.UserResponse, error) {
	transaction, err := u.Repo.UserPreview(req)
	if err != nil {
		return nil, err
	}
	// Fetch file information
	var files []dtof.UploadResponse
	// for _, file := range transaction.File {
	// 	files = append(files, dtof.UploadResponse{
	// 		FileName: file.FileName,
	// 		FilePath: file.FilePath,
	// 	})
	// }

	// roles := dtor.RoleUserResponse{
	// 	RoleID: transaction.Role.RoleID,
	// }

	return &dto.UserResponse{
		ID:          transaction.ID,
		UserID:      transaction.UserID,
		FullName:    transaction.FullName,
		Age:         transaction.Age,
		Description: transaction.Description,
		// Roles:       roles,
		Files: files, // Menambahkan informasi file ke respons

	}, nil
}

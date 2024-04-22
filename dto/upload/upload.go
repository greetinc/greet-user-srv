package dto

type ParamRequest struct {
	FileName string `param:"file_name" validate:"required"`
}

type UploadRequest struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
}

type PictureRequest struct {
	ID       string `json:"id"`
	UserID   string `query:"user_id" validate:"required,user_id"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
}

type UploadResponse struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
}

package dto

import (
	dto "greet-user-srv/dto/limit"
	dtof "greet-user-srv/dto/upload"
)

type UserRequest struct {
	ID string `param:"id" validate:"required"`
}

type UserPreviewRequest struct {
	UserID string `json:"user_id"`
}

type UpdateUserRequest struct {
	ID          string `param:"id" validate:"required"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
}

type UpdateUserResponse struct {
	FullName    string `json:"full_name"`
	Description string `json:"description"`
}

type UserResponse struct {
	ID           string                `json:"id"`
	FullName     string                `json:"full_name"`
	UserID       string                `json:"user_id"`
	Age          int                   `json:"age"`
	Username     string                `json:"username"`
	Email        string                `json:"email"`
	Description  string                `json:"description"`
	Whatsapp     string                `json:"whatsapp"`
	TempatLahir  string                `json:"tempat_lahir"`
	TanggalLahir string                `json:"tanggal_lahir"`
	Gaji         string                `json:"gaji"`
	KTP          string                `json:"ktp"`
	Limits       []dto.LimitResponse   `json:"limits"`
	Files        []dtof.UploadResponse `json:"files"`
}

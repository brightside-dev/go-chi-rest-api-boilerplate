package dtos

type LoginResponseDTO struct {
	User  UserResponseDTO `json:"user"`
	Token string          `json:"token"`
}

package dtos

type LoginResponseDTO struct {
	User  UserDTO `json:"user"`
	Token string  `json:"token"`
}

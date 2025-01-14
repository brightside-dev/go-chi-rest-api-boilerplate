package dtos

import "github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/models"

type UserResponseDTO struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
}

type UsersResponseDTO struct {
	Users *[]models.User `json:"users"`
}

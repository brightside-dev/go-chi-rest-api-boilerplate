package dtos

import "github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/models"

type UserDTO struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
}

type UsersDTO struct {
	Users *[]models.User `json:"users"`
}

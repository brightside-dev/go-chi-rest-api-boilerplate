package services

import (
	"context"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/dtos"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/repositories"

	"github.com/go-chi/jwtauth/v5"
)

type AuthService struct {
	TokenAuth      *jwtauth.JWTAuth
	UserRepository *repositories.UserRepository
}

func NewAuthService(userRepository *repositories.UserRepository,
) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (as *AuthService) Login(
	ctx context.Context,
	email string,
	password string,
) (dtos.LoginResponseDTO, error) {
	// validate username and password and retrieve user from DB
	user := "TODO"

	// generate JWT token
	token := "TODO"

	loginResponseDTO := dtos.LoginResponseDTO{
		User:  user,
		Token: token,
	}

	return loginResponseDTO, nil

}

func (as *AuthService) Logout() {
}

func (as *AuthService) Register() {
}

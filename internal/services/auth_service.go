package services

import (
	"context"
	"strconv"
	"time"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/dtos"
	customErr "github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/errors"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/models"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/repositories"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-chi/jwtauth/v5"
)

type AuthService struct {
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
	// Retrieve the user from the database by email
	users, err := as.UserRepository.FindBy(ctx, "email", email, 0, 0)
	if err != nil {
		return dtos.LoginResponseDTO{}, customErr.ErrFailedToRetrieveUser
	}
	if len(users) == 0 {
		return dtos.LoginResponseDTO{}, customErr.ErrInvalidEmailOrPassword
	}

	user := users[0]

	// Compare the provided password with the hashed password in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return dtos.LoginResponseDTO{}, customErr.ErrInvalidEmailOrPassword
	}

	// Generate JWT token
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{
		"sub": strconv.Itoa(user.ID),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	if err != nil {
		return dtos.LoginResponseDTO{}, err
	}

	// Build response DTO
	userResponseDTO := dtos.UserResponseDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
	}

	responseDTO := dtos.LoginResponseDTO{
		User:  userResponseDTO,
		Token: tokenString,
	}

	return responseDTO, nil

}

func (as *AuthService) Logout() {
}

func (as *AuthService) Register(ctx context.Context, user models.User) (dtos.RegisterResponseDTO, error) {
	// Check if the user already exists
	existingUsers, err := as.UserRepository.FindBy(ctx, "email", user.Email, 0, 0)
	if err != nil {
		return dtos.RegisterResponseDTO{}, customErr.ErrInternalServerError
	}

	if len(existingUsers) > 0 {
		return dtos.RegisterResponseDTO{}, customErr.ErrEmailAlreadyRegistered
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return dtos.RegisterResponseDTO{}, customErr.ErrInternalServerError
	}

	// Create a new user
	newUser := models.User{
		Email:     user.Email,
		Password:  string(hashedPassword),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Country:   user.Country,
		Birthday:  user.Birthday,
	}

	// Save the user to the database
	_, err = as.UserRepository.Insert(ctx, newUser)
	if err != nil {
		return dtos.RegisterResponseDTO{}, err
	}

	responseDTO := dtos.RegisterResponseDTO{
		Message: "user registered successfully",
	}

	return responseDTO, nil
}

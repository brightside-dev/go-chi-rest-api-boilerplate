package validators

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type AuthRequestValidator struct {
	validate *validator.Validate
}

type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=5,max=32,alphanum"`
	Country   string `json:"country" validate:"required"`
	Birthday  string `json:"birthday" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LogoutRequest struct {
	UserID   int64  `json:"user_id" validate:"required"`
	JWTToken string `json:"jwt_token" validate:"required"`
}

func NewAuthRequestValidator() *AuthRequestValidator {
	return &AuthRequestValidator{
		validate: validator.New(),
	}
}

func (ar *AuthRequestValidator) ValidateRequest(req interface{}) error {
	if err := ar.validate.Struct(req); err != nil {
		// Generate user-friendly error messages
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, ar.formatErrorMessage(err))
		}
		return fmt.Errorf("%s", strings.Join(errorMessages, ", "))
	}
	return nil
}

// Format error messages
func (ar *AuthRequestValidator) formatErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", err.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters long", err.Field(), err.Param())
	default:
		return fmt.Sprintf("%s is invalid", err.Field())
	}
}

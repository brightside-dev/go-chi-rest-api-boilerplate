package validators

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type UserRequestValidator struct {
	validate *validator.Validate
}

type CreateUserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Country   string `json:"country" validate:"required"`
	Birthday  string `json:"birthday" validate:"required"`
}
type UpdateUserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Country   string `json:"country" validate:"required"`
	Birthday  string `json:"birthday" validate:"required"`
}

func NewUserRequestValidator() *UserRequestValidator {
	return &UserRequestValidator{
		validate: validator.New(),
	}
}

func (ur *UserRequestValidator) ValidateRequest(req interface{}) error {
	if err := ur.validate.Struct(req); err != nil {
		// Generate user-friendly error messages
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, ur.formatErrorMessage(err))
		}
		return fmt.Errorf("%s", strings.Join(errorMessages, ", "))
	}
	return nil
}

// Format error messages
func (ur *UserRequestValidator) formatErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", err.Field())
	default:
		return fmt.Sprintf("%s is invalid", err.Field())
	}
}

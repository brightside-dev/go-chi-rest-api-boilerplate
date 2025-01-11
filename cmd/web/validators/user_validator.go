package validators

import (
	"errors"
	"regexp"
)

type UserRequestValidator struct {
}

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Country   string `json:"country"`
	Birthday  string `json:"birthday"`
}

func (v *UserRequestValidator) ValidateCreateUserRequest(req CreateUserRequest) error {
	var errorString string

	if req.FirstName == "" {
		errorString += "First name is required. "
	}
	if req.LastName == "" {
		errorString += "Last name is required. "
	}
	if req.Email == "" {
		errorString += "Email is required. "
	}

	if req.Country == "" {
		errorString += "Country is required. "
	}

	// Validate email format using regex
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(emailRegex).MatchString(req.Email) {
		errorString += "Invalid email format. "
	}

	if req.Birthday == "" {
		errorString += "Birthday is required. "
	}

	if errorString != "" {
		return errors.New(errorString)
	}

	return nil
}

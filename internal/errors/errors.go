package errors

import "errors"

var (
	// Generic
	ErrInternalServerError = errors.New("internal server error")

	// Common
	ErrInvalidBody           = errors.New("invalid request body")
	ErrInvalidBirthdayFormat = errors.New("invalid birthday format, expected YYYY-MM-DD")

	// User
	ErrFailedToRetrieveUser = errors.New("failed to retrieve user")
	ErrFailedToInsertUser   = errors.New("failed to insert user")
	ErrFailedToUpdateUser   = errors.New("failed to update user")
	ErrFailedToDeleteUser   = errors.New("failed to delete user")
	ErrInvalidUserID        = errors.New("invalid user ID")

	// Auth
	ErrInvalidJWTToken          = errors.New("invalid JWT token")
	ErrJWTTokenExpired          = errors.New("JWT token has expired")
	ErrFailedToGenerateJWTToken = errors.New("failed to generate JWT token")
	ErrInvalidEmailOrPassword   = errors.New("invalid email or password")
	ErrEmailAlreadyRegistered   = errors.New("email is already registered")
)

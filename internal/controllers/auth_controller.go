package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/models"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/services"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/utils"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/validators"
)

type AuthController struct {
	AuthService *services.AuthService
	Validator   *validators.AuthRequestValidator
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
		Validator:   validators.NewAuthRequestValidator(),
	}
}

func (ac *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	req := validators.RegisterRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		// TODO - refactor custom error message
		err = errors.New("invalid request body")
		utils.WriteAPIErrorResponse(w, r, err)
		return
	}

	if err := ac.Validator.ValidateRequest(req); err != nil {
		utils.WriteAPIErrorResponse(w, r, err)
		return
	}

	// Parse the Birthday string into time.Time
	birthday, err := time.Parse("2006-01-02", req.Birthday)
	if err != nil {
		error := errors.New("invalid birthday format, expected YYYY-MM-DD")
		utils.WriteAPIErrorResponse(w, r, error)
		return
	}

	user := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Country:   req.Country,
		Birthday:  birthday,
	}

	registerResponseDTO, err := ac.AuthService.Register(r.Context(), user)
	if err != nil {
		utils.WriteAPIErrorResponse(w, r, err)
		return
	}

	utils.WriteAPISuccessResponse(w, r, utils.APIResponse{
		Success: true,
		Data:    registerResponseDTO,
	})
}

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	req := validators.LoginRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		// TODO - refactor custom error message
		err = errors.New("invalid request body")
		utils.WriteAPIErrorResponse(w, r, err)
		return
	}

	if err := ac.Validator.ValidateRequest(req); err != nil {
		utils.WriteAPIErrorResponse(w, r, err)
		return
	}

	loginResponseDTO, err := ac.AuthService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		utils.WriteAPIErrorResponse(w, r, err)
		return
	}

	utils.WriteAPISuccessResponse(w, r, utils.APIResponse{
		Success: true,
		Data:    loginResponseDTO,
	})
}

func (ac *AuthController) Logout() {
}

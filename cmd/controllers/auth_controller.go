package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/cmd/utils"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/cmd/validators"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/services"
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

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	req := validators.LoginRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
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

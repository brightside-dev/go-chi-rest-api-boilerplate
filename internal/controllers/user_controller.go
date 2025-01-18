package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	customErr "github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/errors"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/models"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/services"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/utils"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/validators"
	"github.com/go-chi/chi/v5"
)

// DI
type UserController struct {
	UserService *services.UserService
	Validator   *validators.UserRequestValidator
}

// DI
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
		Validator:   validators.NewUserRequestValidator(),
	}
}

func (uc *UserController) Get(w http.ResponseWriter, r *http.Request) {
	// check if the id exists in the URL
	if chi.URLParam(r, "id") == "" {
		utils.WriteAPIErrorResponse(w, r, customErr.ErrInvalidUserID)
		return
	}

	// Get the ID from the URL
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.WriteAPIErrorResponse(w, r, err)
		return
	}

	// Call the user service's Get() method
	userDTO, err := uc.UserService.Get(r.Context(), id)
	if err != nil {
		utils.WriteAPIErrorResponse(w, r, err)
		return
	}

	// Return the successful response
	utils.WriteAPISuccessResponse(w, r, utils.APIResponse{
		Success: true,
		Data:    userDTO,
	})
}

func (uc *UserController) List(w http.ResponseWriter, r *http.Request) {
	// Call the UserRepository's List() method
	usersDTO, err := uc.UserService.List(r.Context())
	if err != nil {
		utils.WriteAPIErrorResponse(w, r, err)
		return
	}

	// Return the successful response
	utils.WriteAPISuccessResponse(w, r, utils.APIResponse{
		Success: true,
		Data:    usersDTO,
	})
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	// initialize the create user request struct
	req := validators.CreateUserRequest{}

	// Decode the request body into the CreateUserRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		utils.WriteAPIErrorResponse(w, r, err)
		return
	}

	// validate the request
	if err := uc.Validator.ValidateRequest(req); err != nil {
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

	// Set the user struct
	user := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Birthday:  birthday,
	}

	// Call the user service's Create() method which returns a userDTO
	userDTO, err := uc.UserService.Create(r.Context(), user)
	if err != nil {
		utils.WriteAPIErrorResponse(w, r, err)
		return
	}

	// Return the successful response
	utils.WriteAPISuccessResponse(w, r, utils.APIResponse{
		Success: true,
		Data:    userDTO,
	})
}

func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) {
	// initialize the update user request struct
	req := validators.UpdateUserRequest{}

	// Decode the request body into the UpdateUserRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		utils.WriteAPIErrorResponse(w, r, err)
		return
	}

	// validate the request
	if err := uc.Validator.ValidateRequest(req); err != nil {
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

	// Get the ID from the URL
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.WriteAPIErrorResponse(w, r, err)
		return
	}

	// Set the user struct
	user := models.User{
		ID:        id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Birthday:  birthday,
	}

	// Call the user service's Update() method which returns a userDTO
	userDTO, err := uc.UserService.Update(r.Context(), user)
	if err != nil {
		utils.WriteAPIErrorResponse(w, r, err)
		return
	}

	// Return the successful response
	utils.WriteAPISuccessResponse(w, r, utils.APIResponse{
		Success: true,
		Data:    userDTO,
	})
}

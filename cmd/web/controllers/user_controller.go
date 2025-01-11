package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/cmd/web/utils"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/cmd/web/validators"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/models"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/services"
	"github.com/go-chi/chi/v5"
)

type UserController struct {
	DB   *sql.DB
	User models.User
}

func (c *UserController) NewUserService(db *sql.DB) *services.UserService {
	return &services.UserService{DB: db}
}

func (uc *UserController) Get(w http.ResponseWriter, r *http.Request) {
	// Create a new user service
	userService := uc.NewUserService(uc.DB)

	// check if the id exists in the URL
	if chi.URLParam(r, "id") == "" {
		utils.WriteAPIErrorResponse(w, r, errors.New("invalid user ID"))
		return
	}

	// Get the ID from the URL
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.WriteAPIErrorResponse(w, r, err)
		return
	}

	// Call the user service's Get() method
	userDTO, err := userService.Get(r.Context(), id)
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
	// Create a new user service
	userService := uc.NewUserService(uc.DB)

	// Call the UserRepository's List() method
	usersDTO, err := userService.List(r.Context())
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
	// initialize the user request validator
	v := validators.UserRequestValidator{}

	// initialize the create user request struct
	req := validators.CreateUserRequest{}

	// Decode the request body into a create user struct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		utils.WriteAPIErrorResponse(w, r, err)
		return
	}

	// Validate the request fields
	if err := v.ValidateCreateUserRequest(req); err != nil {
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
	uc.User = models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Birthday:  birthday,
	}

	// Create a new user service
	userService := uc.NewUserService(uc.DB)

	// Call the user service's Create() method which returns a userDTO
	userDTO, err := userService.Create(r.Context(), uc.User)
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

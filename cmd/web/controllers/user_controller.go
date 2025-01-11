package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/cmd/web/utils"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/models"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/services"
	"github.com/go-chi/chi/v5"
)

type UserController struct {
	DB        *sql.DB
	User      models.User
	Validator *utils.Validator
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
	// parse the request body
	err := json.NewDecoder(r.Body).Decode(&uc.User)
	if err != nil {
		utils.WriteAPIErrorResponse(w, r, err)
		return
	}

	// Close the request body
	defer r.Body.Close()

	// Validate the request body
	validator := &utils.Validator{}
	validator.CheckField(utils.NotBlank(uc.User.FirstName), "first_name", "First name must not be blank")
	validator.CheckField(utils.NotBlank(uc.User.LastName), "last_name", "Last name must not be blank")
	validator.CheckField(utils.NotBlank(uc.User.Email), "email", "Email must not be blank")
	validator.CheckField(utils.NotBlank(uc.User.Birthday), "birthday", "Birthday must not be blank")
	validator.CheckField(utils.NotBlank(uc.User.Country), "country", "Country must not be blank")

	if !validator.Valid() {
		errorsStruct := utils.Errors{}
		errorsStruct.Errors = validator.FieldErrors

		errorsString, err := json.Marshal(errorsStruct)
		if err != nil {
			utils.WriteAPIErrorResponse(w, r, err)
			return
		}

		err = errors.New(string(errorsString))
		utils.WriteAPIErrorResponse(w, r, err)
		return
	}

	// Create a new user service
	userService := uc.NewUserService(uc.DB)

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

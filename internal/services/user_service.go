package services

import (
	"context"
	"database/sql"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/dtos"
	customErr "github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/errors"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/models"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/repositories"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func (us *UserService) NewUserRepository(db *sql.DB) *repositories.UserRepository {
	return &repositories.UserRepository{DB: db}
}

func (us *UserService) Get(ctx context.Context, id int) (dtos.UserResponseDTO, error) {
	user, err := us.UserRepository.FindOneById(ctx, id)
	if err != nil {
		return dtos.UserResponseDTO{}, customErr.ErrFailedToRetrieveUser
	}

	dto := dtos.UserResponseDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
	}

	return dto, nil
}

func (us *UserService) List(ctx context.Context) (dtos.UsersResponseDTO, error) {
	users, err := us.UserRepository.FindAll(ctx, 0, 0)
	if err != nil {
		return dtos.UsersResponseDTO{}, customErr.ErrFailedToRetrieveUser
	}

	dto := dtos.UsersResponseDTO{
		Users: &users,
	}

	return dto, nil
}

func (us *UserService) Create(ctx context.Context, user models.User) (dtos.UserResponseDTO, error) {
	dbUser, err := us.UserRepository.FindBy(ctx, "email", user.Email, 0, 0)
	if err != nil {
		return dtos.UserResponseDTO{}, err
	}

	if dbUser != nil {
		return dtos.UserResponseDTO{}, customErr.ErrEmailAlreadyRegistered
	}

	newUser, err := us.UserRepository.Insert(ctx, user)
	if err != nil {
		return dtos.UserResponseDTO{}, customErr.ErrFailedToInsertUser
	}

	dto := dtos.UserResponseDTO{
		ID:        newUser.ID,
		FirstName: newUser.FirstName,
	}

	return dto, nil
}

func (us *UserService) Update(ctx context.Context, user models.User) (dtos.UserResponseDTO, error) {
	dto := dtos.UserResponseDTO{}

	dbUser, err := us.UserRepository.FindOneById(ctx, user.ID)
	if err != nil {
		return dtos.UserResponseDTO{}, err
	}

	// check dbUser struct is the same as user struct
	if dbUser.FirstName == user.FirstName &&
		dbUser.LastName == user.LastName &&
		dbUser.Email == user.Email &&
		dbUser.Birthday == user.Birthday &&
		dbUser.Country == user.Country {
		// No changes, so no need to update
		dto.ID = dbUser.ID
		dto.FirstName = dbUser.FirstName

		return dto, nil
	} else {
		updatedUser, err := us.UserRepository.Update(ctx, user)
		if err != nil {
			return dtos.UserResponseDTO{}, customErr.ErrFailedToUpdateUser
		}

		dto.ID = updatedUser.ID
		dto.FirstName = updatedUser.FirstName
	}

	return dto, nil
}

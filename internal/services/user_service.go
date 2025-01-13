package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/dtos"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/models"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/repositories"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func (us *UserService) NewUserRepository(db *sql.DB) *repositories.UserRepository {
	return &repositories.UserRepository{DB: db}
}

func (us *UserService) Get(ctx context.Context, id int) (dtos.UserDTO, error) {
	user, err := us.UserRepository.FindById(ctx, id)
	if err != nil {
		return dtos.UserDTO{}, err
	}

	userDTO := dtos.UserDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
	}

	return userDTO, nil
}

func (us *UserService) List(ctx context.Context) (dtos.UsersDTO, error) {
	users, err := us.UserRepository.FindAll(ctx, 0, 0)
	if err != nil {
		return dtos.UsersDTO{}, err
	}

	dto := dtos.UsersDTO{
		Users: &users,
	}

	return dto, nil
}

func (us *UserService) Create(ctx context.Context, user models.User) (dtos.UserDTO, error) {
	dbUser, err := us.UserRepository.FindBy(ctx, "email", user.Email, 0, 0)
	if err != nil {
		return dtos.UserDTO{}, err
	}

	if dbUser != nil {
		return dtos.UserDTO{}, errors.New("user with email already exists")
	}

	newUser, err := us.UserRepository.Insert(ctx, user)
	if err != nil {
		return dtos.UserDTO{}, err
	}

	dto := dtos.UserDTO{
		ID:        newUser.ID,
		FirstName: newUser.FirstName,
	}

	return dto, nil
}

func (us *UserService) Update(ctx context.Context, user models.User) (dtos.UserDTO, error) {
	dto := dtos.UserDTO{}

	dbUser, err := us.UserRepository.FindById(ctx, user.ID)
	if err != nil {
		return dtos.UserDTO{}, err
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
			return dtos.UserDTO{}, err
		}

		dto.ID = updatedUser.ID
		dto.FirstName = updatedUser.FirstName
	}

	return dto, nil
}

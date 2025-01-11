package services

import (
	"context"
	"database/sql"

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

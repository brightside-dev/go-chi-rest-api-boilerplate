package container

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/repositories"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/services"
)

type Container struct {
	DB             *sql.DB
	UserService    *services.UserService
	UserRepository *repositories.UserRepository
	Logger         *slog.Logger
}

func NewContainer(db *sql.DB) *Container {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	userRepository := &repositories.UserRepository{DB: db}
	userService := &services.UserService{UserRepository: userRepository}

	return &Container{
		DB:             db,
		UserService:    userService,
		UserRepository: userRepository,
		Logger:         logger,
	}
}

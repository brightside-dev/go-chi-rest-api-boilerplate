package web

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/database"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/repositories"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Container struct {
	DB             *sql.DB
	UserService    *services.UserService
	UserRepository *repositories.UserRepository
	Logger         *slog.Logger
}

func StartServer() {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	// To keep the main() function tidy I've put the code for creating a connection
	// pool into the separate openDB() function below. We pass openDB() the DSN
	// from the command-line flag.
	db, err := database.OpenMySQLDB()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// We also defer a call to db.Close(), so that the connection pool is closed
	// before the main() function exits.
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	container := newContainer(db, logger)

	// Setup routes with the DI container
	SetupRoutes(r, container)

	http.ListenAndServe(":3000", r)
}

func newContainer(db *sql.DB, logger *slog.Logger) *Container {
	userRepository := &repositories.UserRepository{DB: db}
	userService := &services.UserService{UserRepository: userRepository}

	return &Container{
		DB:             db,
		UserService:    userService,
		UserRepository: userRepository,
		Logger:         logger,
	}
}

package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/repositories"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

type Container struct {
	Config         *Config
	DB             *sql.DB
	UserService    *services.UserService
	UserRepository *repositories.UserRepository
	AuthService    *services.AuthService
	Logger         *slog.Logger
}

type Config struct {
	DBUser    string
	DBPass    string
	DBHost    string
	DBPort    string
	DBName    string
	JWTSecret string
}

func main() {
	// Initialize Chi router
	r := chi.NewRouter()

	// Middleware chain
	r.Use(middleware.Recoverer)

	// Initialize logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	// Initialize config
	config := newConfig(logger)

	// Initialize DB connection
	db, err := openDBConnection(*config, logger)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// Defer closing the DB connection
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			os.Exit(1)
		}
	}(db)

	// Initialize DI container
	container := newContainer(config, db, logger)

	// Initialize routes
	SetupRoutes(r, container)

	// Start the server
	http.ListenAndServe(":3000", r)
}

func newContainer(config *Config, db *sql.DB, logger *slog.Logger) *Container {
	userRepository := &repositories.UserRepository{DB: db}
	userService := &services.UserService{UserRepository: userRepository}
	authService := &services.AuthService{UserRepository: userRepository}

	return &Container{
		Config:         config,
		DB:             db,
		UserService:    userService,
		AuthService:    authService,
		UserRepository: userRepository,
		Logger:         logger,
	}
}

func newConfig(logger *slog.Logger) *Config {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
		os.Exit(1)
	}

	config := &Config{DBUser: os.Getenv("DB_USER"),
		DBPass:    os.Getenv("DB_PASSWORD"),
		DBHost:    os.Getenv("DB_HOST"),
		DBPort:    os.Getenv("DB_PORT"),
		DBName:    os.Getenv("DB_NAME"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	// Check that none of the config values are empty
	if config.DBUser == "" ||
		config.DBPass == "" ||
		config.DBHost == "" ||
		config.DBPort == "" ||
		config.DBName == "" ||
		config.JWTSecret == "" {
		logger.Error("One or more required environment variables are missing")
		os.Exit(1)
	}

	return config
}

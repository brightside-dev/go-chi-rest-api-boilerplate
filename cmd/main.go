package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/config"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/db"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/repositories"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/routes"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
)

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
	conf := newConfig(logger)

	// Initialize DB connection
	db, err := db.OpenDB(conf, logger)
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
	container := newContainer(*conf, db, logger)

	// Initialize routes
	routes.SetupRoutes(r, container)

	// Start the server
	http.ListenAndServe(":3000", r)
}

func newContainer(conf config.Config, db *sql.DB, logger *slog.Logger) *config.Container {
	userRepository := &repositories.UserRepository{DB: db}
	userService := &services.UserService{UserRepository: userRepository}
	authService := &services.AuthService{UserRepository: userRepository}
	tokenAuth := jwtauth.New("HS256", []byte(conf.JWTSecret), nil)

	return &config.Container{
		Config:         &conf,
		DB:             db,
		UserService:    userService,
		AuthService:    authService,
		UserRepository: userRepository,
		Logger:         logger,
		TokenAuth:      tokenAuth,
	}
}

func newConfig(logger *slog.Logger) *config.Config {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
		os.Exit(1)
	}

	config := &config.Config{DBUser: os.Getenv("DB_USER"),
		DBPass:        os.Getenv("DB_PASSWORD"),
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBName:        os.Getenv("DB_NAME"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		MailgunDomain: os.Getenv("MAILGUN_DOMAIN"),
		MailgunAPIKey: os.Getenv("MAILGUN_API_KEY"),
	}

	// Check that none of the config values are empty
	if config.DBUser == "" ||
		config.DBPass == "" ||
		config.DBHost == "" ||
		config.DBPort == "" ||
		config.DBName == "" ||
		config.JWTSecret == "" ||
		config.MailgunDomain == "" ||
		config.MailgunAPIKey == "" {
		logger.Error("One or more required environment variables are missing")
		os.Exit(1)
	}

	return config
}

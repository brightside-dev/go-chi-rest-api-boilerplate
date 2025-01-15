package config

import (
	"database/sql"
	"html/template"
	"log/slog"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/repositories"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/services"
	"github.com/go-chi/jwtauth/v5"
)

type Container struct {
	Config         *Config
	DB             *sql.DB
	TemplateCache  map[string]*template.Template
	UserService    *services.UserService
	UserRepository *repositories.UserRepository
	AuthService    *services.AuthService
	Logger         *slog.Logger
	TokenAuth      *jwtauth.JWTAuth
}

type Config struct {
	DBUser        string
	DBPass        string
	DBHost        string
	DBPort        string
	DBName        string
	JWTSecret     string
	MailgunDomain string
	MailgunAPIKey string
	FirebaseKey   string
}

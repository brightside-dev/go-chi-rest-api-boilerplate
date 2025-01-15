package routes

import (
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/config"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/controllers"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func SetupRoutes(r *chi.Mux, container *config.Container) {
	// Create a new authController
	authController := controllers.NewAuthController(container.AuthService)

	// rate APIs
	r.Group(func(r chi.Router) {
		// Use the JWT token verifier middleware
		r.Use(jwtauth.Verifier(container.TokenAuth))
		r.Use(middlewares.Auth)

		// Create a new userController
		userController := controllers.NewUserController(container.UserService)

		r.Get("/api/users", userController.List)
		r.Post("/api/users", userController.Create)
		r.Get("/api/users/{id}", userController.Get)
		r.Put("/api/users/update/{id}", userController.Update)

		r.Post("/api/auth/register", authController.Register)
	})

	// Public APIs
	r.Group(func(r chi.Router) {
		r.Post("/api/auth/login", authController.Login)
		r.Post("/api/auth/register", authController.Register)
	})
}

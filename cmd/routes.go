package main

import (
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/cmd/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func SetupRoutes(r *chi.Mux, container *Container) {
	// Create a new authController
	authController := controllers.NewAuthController(container.AuthService)

	// rate APIs
	r.Group(func(r chi.Router) {
		// Use the JWT token verifier middleware
		r.Use(jwtauth.Verifier(container.TokenAuth))
		r.Use(JWTAuthMiddleware)

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

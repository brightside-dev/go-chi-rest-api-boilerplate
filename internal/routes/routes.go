package routes

import (
	"net/http"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/config"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/controllers"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/middlewares"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/templates"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func SetupRoutes(r *chi.Mux, container *config.Container) {
	// Create a new authController
	authController := controllers.NewAuthController(container.AuthService)
	// Create a new userController
	userController := controllers.NewUserController(container.UserService)

	// Private APIs
	r.Group(func(r chi.Router) {
		// Use the JWT token verifier middleware
		r.Use(jwtauth.Verifier(container.TokenAuth))
		// Use the custom Auth middleware
		r.Use(middlewares.Auth)

		r.Get("/api/users", userController.List)
		r.Post("/api/users", userController.Create)
		r.Get("/api/users/{id}", userController.Get)
		r.Put("/api/users/update/{id}", userController.Update)
	})

	// Public APIs
	r.Group(func(r chi.Router) {
		r.Post("/api/auth/login", authController.Login)
		r.Post("/api/auth/register", authController.Register)
	})

	// Serve the static files
	r.Get("/*", http.FileServer(http.Dir("public")).ServeHTTP)
	// Admin Views
	r.Group(func(r chi.Router) {
		r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
			templates.Render(w, r, "index", nil, container.TemplateCache)
		})
	})
}

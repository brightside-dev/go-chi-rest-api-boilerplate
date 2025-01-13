package web

import (
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/cmd/web/controllers"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux, container *Container) {
	// Public APIs
	r.Group(func(r chi.Router) {
		// Create a new userController
		userController := controllers.NewUserController(container.UserService)

		r.Get("/api/users", userController.List)
		r.Post("/api/users", userController.Create)
		r.Get("/api/users/{id}", userController.Get)
		r.Put("/api/users/update/{id}", userController.Update)
	})
}

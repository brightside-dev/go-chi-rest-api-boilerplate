package web

import (
	"database/sql"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/cmd/web/controllers"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux, db *sql.DB) {
	// Public APIs
	r.Group(func(r chi.Router) {
		// Create a new userController
		userController := controllers.UserController{DB: db}

		r.Get("/api/users", userController.List)
		r.Post("/api/users", userController.Create)
		r.Get("/api/users/{id}", userController.Get)
	})
}

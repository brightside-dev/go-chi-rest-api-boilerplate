package web

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/database"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/container"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

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

	// Initialize the DI container
	container := container.NewContainer(db)

	// Setup routes with the DI container
	SetupRoutes(r, container)

	http.ListenAndServe(":3000", r)
}

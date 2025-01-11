package database

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool
// for a given DSN
func OpenMySQLDB() (*sql.DB, error) {
	// Init logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
		return nil, err
	}

	// Get database connection details from environment variables
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		logger.Error("DB_USER environment variable is not set")
		return nil, errors.New("DB_USER environment variable is not set")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		logger.Error("DB_PASSWORD environment variable is not set")
		return nil, errors.New("DB_PASSWORD environment variable is not set")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		logger.Error("DB_HOST environment variable is not set")
		return nil, errors.New("DB_HOST environment variable is not set")
	}

	// Create a DSN from the environment variables
	dsn := dbUser + ":" + dbPassword + "@/" + dbHost

	// Open a database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	// Ping the database to check if the connection is working
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return db, nil
}

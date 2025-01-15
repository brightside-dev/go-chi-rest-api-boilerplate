package db

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool
// for a given DSN
func OpenDB(config *config.Config, logger *slog.Logger) (*sql.DB, error) {

	// Create a DSN from the environment variables
	dsn := config.DBUser + ":" + config.DBPass + "@/" + config.DBHost

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

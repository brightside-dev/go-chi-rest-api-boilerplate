package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"time"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/models"
)

var SearchableFields = map[string]FieldMeta{
	"id":         {Allowed: true, Type: "int"},
	"first_name": {Allowed: true, Type: "string"},
	"last_name":  {Allowed: true, Type: "string"},
	"email":      {Allowed: true, Type: "string"},
}

type UserRepository struct {
	DB *sql.DB
}

func (rp *UserRepository) GetSearchableFields() map[string]FieldMeta {
	return SearchableFields
}

func parseBirthday(birthday interface{}) (time.Time, error) {
	switch v := birthday.(type) {
	case string:
		return time.Parse("2006-01-02", v)
	case []byte:
		return time.Parse("2006-01-02", string(v))
	default:
		return time.Time{}, fmt.Errorf("unexpected type for birthday: %T", v)
	}
}

func (rp *UserRepository) scanRow(r *sql.Row, u *models.User) error {
	var birthdayRaw interface{}
	// Scan the values from the row
	err := r.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&birthdayRaw, // Scan the raw value of birthday first
		&u.Country,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return err
	}

	// Handle the birthday conversion
	u.Birthday, err = parseBirthday(birthdayRaw)
	if err != nil {
		return fmt.Errorf("invalid birthday format: %v", err)
	}

	return nil
}

func (rp *UserRepository) scanRows(r *sql.Rows, u *models.User) error {
	var birthdayRaw interface{}
	// Scan the values from the row
	err := r.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&birthdayRaw, // Scan the raw value of birthday first
		&u.Country,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return err
	}

	// Handle the birthday conversion
	u.Birthday, err = parseBirthday(birthdayRaw)
	if err != nil {
		return fmt.Errorf("invalid birthday format: %v", err)
	}

	return nil
}

// Standard repository methods
func (rp *UserRepository) Insert(ctx context.Context, u models.User) (models.User, error) {
	// Begin a transaction
	tx, err := rp.DB.BeginTx(ctx, nil)
	if err != nil {
		return models.User{}, err
	}

	// Ensure rollback is called if the function exits without committing
	defer func() {
		tx.Rollback()
	}()

	// Insert statement
	stmt := `INSERT INTO users (first_name, last_name, email, birthday, country) VALUES (?, ?, ?, ?, ?)`

	// Execute the query

	// Using the tx.ExecContext() method to execute the query and get last user id
	result, err := tx.ExecContext(ctx, stmt, u.FirstName, u.LastName, u.Email, u.Birthday, u.Country)
	if err != nil {
		return models.User{}, err
	}

	// Get the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return models.User{}, err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return models.User{}, err
	}

	// Set the ID of the user
	u.ID = int(id)

	return u, nil
}

func (rp *UserRepository) Update(ctx context.Context, u models.User) (int, error) {
	tx, err := rp.DB.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer func() {
		tx.Rollback()
	}()

	// Update statement
	stmt := `UPDATE users SET first_name = ?, last_name = ?, email = ?, birthday = ?, country = ? WHERE id = ?`

	result, err := tx.ExecContext(ctx, stmt, u.FirstName, u.LastName, u.Email, u.Birthday, u.Country, u.ID)
	if err != nil {
		return 0, err
	}

	// Retrieve the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return int(rowsAffected), nil
}

func (rp *UserRepository) Delete(ctx context.Context, id int) (int, error) {
	tx, err := rp.DB.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer func() {
		tx.Rollback()
	}()

	// Delete statement
	stmt := `DELETE FROM users WHERE id = ?`

	result, err := tx.ExecContext(ctx, stmt, id)
	if err != nil {
		return 0, err
	}

	// Retrieve the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return int(rowsAffected), nil
}

func (rp *UserRepository) FindById(ctx context.Context, id int) (models.User, error) {
	// Query to find user by id
	stmt := `SELECT * FROM users WHERE id = ?`

	// Execute the query
	row := rp.DB.QueryRow(stmt, id)

	// create empty user struct
	user := models.User{}

	// Scan the row
	err := rp.scanRow(row, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (rp *UserRepository) FindAll(ctx context.Context, limit int, offset int) ([]models.User, error) {
	// Base query
	stmt := `SELECT * FROM users`
	var args []interface{}

	// Add limit if greater than 0
	if limit > 0 {
		stmt += " LIMIT ?"
		args = append(args, limit)
	}

	// Add offset if greater than 0
	if offset > 0 {
		stmt += " OFFSET ?"
		args = append(args, offset)
	}

	// Execute the query
	rows, err := rp.DB.QueryContext(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create an empty users slice
	users := []models.User{}

	// scan the rows
	for rows.Next() {
		var user models.User
		err := rp.scanRows(rows, &user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (rp *UserRepository) FindBy(ctx context.Context, field string, value interface{}, limit int, offset int) ([]models.User, error) {
	// Check if the field is searchable
	fieldMeta, ok := SearchableFields[field]
	if !ok {
		return nil, fmt.Errorf("field %s is not searchable", field)
	}

	// Check if the value type matches the expected type for the field
	expectedType := fieldMeta.Type
	actualType := reflect.TypeOf(value).String()

	if expectedType != actualType {
		return nil, fmt.Errorf("value type for field %s should be %s, got %s", field, expectedType, actualType)
	}

	// Base query
	stmt := fmt.Sprintf("SELECT * FROM users WHERE %s = ?", field)

	// Arguments slice for query placeholders
	args := []interface{}{value}

	// Add LIMIT and OFFSET if applicable
	if limit > 0 {
		stmt += " LIMIT ?"
		args = append(args, limit)
	}
	if offset > 0 {
		stmt += " OFFSET ?"
		args = append(args, offset)
	}

	// Execute the query
	rows, err := rp.DB.Query(stmt, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create an empty users slice
	users := []models.User{}

	// scan the rows
	for rows.Next() {
		var user models.User
		err := rp.scanRows(rows, &user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Custom repository methods
func (rp *UserRepository) FindUserWithProfile(ctx context.Context, id int) (models.User, error) {
	// create empty user
	user := models.User{}

	return user, nil
}

package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/models"
)

var AdminUserSearchableFields = map[string]FieldMeta{
	"id":         {Allowed: true, Type: "int"},
	"first_name": {Allowed: true, Type: "string"},
	"last_name":  {Allowed: true, Type: "string"},
	"email":      {Allowed: true, Type: "string"},
}

type AdminUserRepository struct {
	DB *sql.DB
}

func (rp *AdminUserRepository) scanRow(r *sql.Row, u *models.AdminUser) error {
	// Scan the values from the row
	err := r.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (rp *AdminUserRepository) scanRows(r *sql.Rows, u *models.AdminUser) error {
	// Scan the values from the row
	err := r.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (rp *AdminUserRepository) Insert(ctx context.Context, u *models.AdminUser) (*models.AdminUser, error) {
	// Start a transaction
	tx, err := rp.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // Rollback if commit is not successful

	// Execute the insert statement
	stmt := `INSERT INTO admin_users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)`
	result, err := tx.ExecContext(ctx, stmt, u.FirstName, u.LastName, u.Email, u.Password)
	if err != nil {
		return nil, err
	}

	// Retrieve the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// Set the ID and return the user
	u.ID = int(id)
	return u, nil
}

func (rp *AdminUserRepository) Update(ctx context.Context, u *models.AdminUser) (*models.AdminUser, error) {
	// Start a transaction
	tx, err := rp.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // Rollback if commit is not successful

	// Execute the update statement
	stmt := `UPDATE admin_users SET first_name = ?, last_name = ?, email = ?, password = ? WHERE id = ?`
	_, err = tx.ExecContext(ctx, stmt, u.FirstName, u.LastName, u.Email, u.Password, u.ID)
	if err != nil {
		return nil, err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return u, nil
}

func (rp *AdminUserRepository) Delete(ctx context.Context, id int) error {
	// Start a transaction
	tx, err := rp.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() // Rollback if commit is not successful

	// Execute the delete statement
	stmt := `DELETE FROM admin_users WHERE id = ?`
	_, err = tx.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (rp *AdminUserRepository) FindOneById(ctx context.Context, id int) (*models.AdminUser, error) {
	stmt := `SELECT * FROM admin_users WHERE id = ?`
	row := rp.DB.QueryRow(stmt, id)

	u := &models.AdminUser{}
	err := rp.scanRow(row, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (rp *AdminUserRepository) FindAll(ctx context.Context, limit int, offset int) ([]models.AdminUser, error) {
	// Base query
	stmt := `SELECT * FROM admin_users`
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
	adminUsers := []models.AdminUser{}

	// scan the rows
	for rows.Next() {
		var adminUser models.AdminUser
		err := rp.scanRows(rows, &adminUser)
		if err != nil {
			return nil, err
		}
		adminUsers = append(adminUsers, adminUser)
	}

	return adminUsers, nil
}

func (rp *AdminUserRepository) FindBy(ctx context.Context, field string, value interface{}, limit int, offset int) ([]models.AdminUser, error) {
	// Check if the field is searchable
	fieldMeta, ok := AdminUserSearchableFields[field]
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
	adminUsers := []models.AdminUser{}

	// scan the rows
	for rows.Next() {
		var adminUser models.AdminUser
		err := rp.scanRows(rows, &adminUser)
		if err != nil {
			return nil, err
		}
		adminUsers = append(adminUsers, adminUser)
	}

	return adminUsers, nil
}

func (rp *AdminUserRepository) FindByOneEmail(ctx context.Context, email string) (*models.AdminUser, error) {
	stmt := `SELECT * FROM admin_users WHERE email = ?"`
	row := rp.DB.QueryRow(stmt, email)

	u := &models.AdminUser{}
	err := rp.scanRow(row, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

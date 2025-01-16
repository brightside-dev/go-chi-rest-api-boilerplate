package models

import (
	"database/sql"
)

const AdminUserTable = "admin_users"

type AdminUser struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AdminUserModel struct {
	DB *sql.DB
}

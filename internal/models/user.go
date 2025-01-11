package models

import (
	"database/sql"
	"time"
)

const UserTable = "users"

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Birthday  time.Time `json:"birthday"`
	Country   string    `json:"country"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

type UserModel struct {
	DB *sql.DB
}

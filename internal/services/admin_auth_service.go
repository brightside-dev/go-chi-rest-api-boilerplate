package services

import (
	"context"
	"database/sql"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/repositories"
)

type AdminAuthService struct {
	AdminUserRepository *repositories.AdminUserRepository
}

func (aa *AdminAuthService) NewAdminUserRepository(db *sql.DB) *repositories.AdminUserRepository {
	return &repositories.AdminUserRepository{DB: db}
}

func (aa *AdminAuthService) Login(ctx context.Context, email, password string) error {
	return nil
}

func (aa *AdminAuthService) Logout(ctx context.Context, id int) error {
	return nil
}

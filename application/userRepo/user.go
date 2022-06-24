package userRepo

import (
	"github.com/Tambarie/evolve-credit/domain/models"
)

// Repository interface
type Repository interface {
	CreateUser(userDto *models.User) (*models.User, error)
	UpdateUser(userDto *models.User) (*models.User, error)
	GetUserByEmail(email string) ([]*models.User, error)
	GetUserByDate(date string) ([]*models.User, error)
}

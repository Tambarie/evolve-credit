package userRepo

import (
	"github.com/Tambarie/evolve-credit/domain/models"
	"time"
)

// Repository interface
type Repository interface {
	CreateUser(userDto *models.User) (*models.User, error)
	UpdateUser(userID string, update *models.UpdateUserInfo) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	GetUserByDate(firstdate, seconddate time.Time) ([]models.User, error)
	GetUserDetailsByEmail(email string) (*models.User, error)
	FindUserByID(id string) (*models.User, error)
	GetAllUsers(limit int) ([]*models.User, error)
}

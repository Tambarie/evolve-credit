package userRepo

import (
	"github.com/Tambarie/evolve-credit/domain/models"
	"gorm.io/gorm"
)

// RepositoryDB struct
type RepositoryDB struct {
	db *gorm.DB
}

// NewUserRepositoryDB function to initialize RepositoryDB struct
func NewUserRepositoryDB(client *gorm.DB) *RepositoryDB {
	return &RepositoryDB{
		db: client,
	}
}

// GetUserByEmail Fetches user based on the email from the database
func (userRepo *RepositoryDB) GetUserByEmail(email string) ([]*models.User, error) {
	return nil, nil

}

// CreateUser  Creates models to the database
func (userRepo *RepositoryDB) CreateUser(u *models.User) (*models.User, error) {
	return nil, nil
}

// UpdateUser  Creates models to the database
func (userRepo *RepositoryDB) UpdateUser(userDto *models.User) (*models.User, error) {
	return nil, nil
}

// GetUserByDate  Fetches user based on the date from the database
func (userRepo *RepositoryDB) GetUserByDate(date string) ([]*models.User, error) {
	return nil, nil
}

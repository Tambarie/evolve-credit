package userRepo

import (
	"errors"
	"github.com/Tambarie/evolve-credit/domain/models"
	"gorm.io/gorm"
	"log"
	"time"
)

// RepositoryDB struct
type RepositoryDB struct {
	DB *gorm.DB
}

// NewUserRepositoryDB function to initialize RepositoryDB struct
func NewUserRepositoryDB(client *gorm.DB) *RepositoryDB {
	return &RepositoryDB{
		DB: client,
	}
}

// FindUserByEmail Fetches user based on the email from the database
func (userRepo *RepositoryDB) FindUserByEmail(email string) (*models.User, error) {
	result := userRepo.DB.First(&models.User{}, "email = ?", email)
	if result.RowsAffected > 0 {
		return nil, errors.New("user already exists")
	}
	return nil, nil
}

// CreateUser  Creates models to the database
func (userRepo *RepositoryDB) CreateUser(user *models.User) (*models.User, error) {
	err := userRepo.DB.Create(user).Error
	return nil, err
}

// UpdateUser  Creates models to the database
func (userRepo *RepositoryDB) UpdateUser(userID string, update *models.UpdateUserInfo) (*models.User, error) {
	user := &models.User{}
	err := userRepo.DB.Model(user).Where("id = ?", userID).Updates(&models.User{

		FirstName:   update.FirstName,
		LastName:    update.LastName,
		DateOfBirth: update.DateOfBirth,
		UpdateAt:    update.UpdateAt,
	}).Error
	return user, err
}

// GetUserByDate  Fetches user based on the date from the database
func (userRepo *RepositoryDB) GetUserByDate(firstdate, seconddate time.Time) ([]models.User, error) {
	log.Println(firstdate, seconddate)
	var users []models.User
	err := userRepo.DB.Where("created_at BETWEEN ? AND ?", firstdate, seconddate).Find(&users).Error
	return users, err

}

func (userRepo *RepositoryDB) GetUserDetailsByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := userRepo.DB.Where("email = ?", email).First(user).Error
	return user, err
}

func (userRepo *RepositoryDB) FindUserByID(id string) (*models.User, error) {
	result := userRepo.DB.First(&models.User{}, "id = ?", id)
	if result.RowsAffected > 0 {
		return nil, errors.New("user already exists")
	}
	return nil, nil
}

func (userRepo RepositoryDB) GetAllUsers(limit int) ([]*models.User, error) {

	if limit == 0 {
		limit = 10
	}

	var users []*models.User
	err := userRepo.DB.Limit(limit).Find(&users).Error
	return users, err
}

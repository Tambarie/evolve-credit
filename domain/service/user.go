package service

import (
	"github.com/Tambarie/evolve-credit/application/userRepo"
	"github.com/Tambarie/evolve-credit/domain/models"
	"time"
)

// UserService interface
type UserService interface {
	FindUserByEmail(email string) (*models.User, error)
	FindUserByID(id string) (*models.User, error)
	CreateUser(userDto *models.User) (*models.User, error)
	UpdateUser(userID string, update *models.UpdateUserInfo) (*models.User, error)
	GetUserByDate(firstdate, seconddate time.Time) ([]models.User, error)
	GetUserDetailsByEmail(email string) (*models.User, error)
	GetAllUsers(limit int) ([]*models.User, error)
}

// UserServiceRepository struct
type UserServiceRepository struct {
	repo userRepo.Repository
}

func NewUserService(repo userRepo.Repository) *UserServiceRepository {
	return &UserServiceRepository{
		repo: repo,
	}
}

func (u *UserServiceRepository) CreateUser(userDto *models.User) (*models.User, error) {
	return u.repo.CreateUser(userDto)
}

func (u *UserServiceRepository) FindUserByEmail(email string) (*models.User, error) {
	return u.repo.FindUserByEmail(email)
}

func (u *UserServiceRepository) UpdateUser(userID string, update *models.UpdateUserInfo) (*models.User, error) {
	return u.repo.UpdateUser(userID, update)
}

func (u *UserServiceRepository) GetUserByDate(firstdate, seconddate time.Time) ([]models.User, error) {
	return u.repo.GetUserByDate(firstdate, seconddate)
}

func (u *UserServiceRepository) GetUserDetailsByEmail(email string) (*models.User, error) {
	return u.repo.GetUserDetailsByEmail(email)
}

func (u *UserServiceRepository) FindUserByID(id string) (*models.User, error) {
	return u.repo.FindUserByID(id)
}

func (u *UserServiceRepository) GetAllUsers(limit int) ([]*models.User, error) {
	return u.repo.GetAllUsers(limit)
}

package service

import (
	"github.com/Tambarie/evolve-credit/application/userRepo"
	"github.com/Tambarie/evolve-credit/domain/models"
)

// UserService interface
type UserService interface {
	CreateUser(userDto *models.User) (*models.User, error)
	UpdateUser(userDto *models.User) (*models.User, error)
	GetUserByEmail(email string) ([]*models.User, error)
	GetUserByDate(date string) ([]*models.User, error)
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

func (u *UserServiceRepository) GetUserByEmail(email string) ([]*models.User, error) {
	return u.repo.GetUserByEmail(email)
}

func (u *UserServiceRepository) UpdateUser(userDto *models.User) (*models.User, error) {
	return u.repo.UpdateUser(userDto)
}

func (u *UserServiceRepository) GetUserByDate(date string) ([]*models.User, error) {
	return u.repo.GetUserByDate(date)
}

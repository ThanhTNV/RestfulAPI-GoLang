package service

import (
	"errors"

	"github.com/ThanhTNV/RestfulAPI-GoLang/internal/models"
	"github.com/ThanhTNV/RestfulAPI-GoLang/internal/repository"
	"gorm.io/gorm"
)

// UserService defines the interface for user business logic
type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(id uint, user *models.User) error
	DeleteUser(id uint) error
}

// userService implements UserService interface
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// GetAllUsers retrieves all users
func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}

// GetUserByID retrieves a user by ID
func (s *userService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

// CreateUser creates a new user
func (s *userService) CreateUser(user *models.User) error {
	// Business logic can be added here (e.g., validation, email uniqueness check)
	return s.userRepo.Create(user)
}

// UpdateUser updates an existing user
func (s *userService) UpdateUser(id uint, updateData *models.User) error {
	// Check if user exists
	existingUser, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	// Update fields
	existingUser.Name = updateData.Name
	existingUser.Email = updateData.Email

	return s.userRepo.Update(existingUser)
}

// DeleteUser deletes a user by ID
func (s *userService) DeleteUser(id uint) error {
	err := s.userRepo.Delete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}
	return nil
}

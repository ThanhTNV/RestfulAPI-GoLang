package repository

import (
	"github.com/ThanhTNV/RestfulAPI-GoLang/internal/models"
	"gorm.io/gorm"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByID(id uint) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id uint) error
}

// userRepository implements UserRepository interface
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// FindAll retrieves all users from the database
func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	result := r.db.Find(&users)
	return users, result.Error
}

// FindByID retrieves a user by ID
func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// Create creates a new user
func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// Update updates an existing user
func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

// Delete deletes a user by ID
func (r *userRepository) Delete(id uint) error {
	result := r.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

package models

// User represents a user entity
type User struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

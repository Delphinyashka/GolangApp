package models

import (
	"auth/config"
	"errors"
)

// User struct represents the user model.
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"-"` // Don't expose hashed password in JSON response
}

// GetUserByUsername retrieves a user by username.
func GetUserByUsername(username string) (*User, error) {
	db := config.GetDB()
	var user User

	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

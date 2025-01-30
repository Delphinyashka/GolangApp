package models

import (
	"auth/config"
	"errors"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"-"`
}

func GetUserByUsername(username string) (*User, error) {
	db := config.GetDB()
	var user User

	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

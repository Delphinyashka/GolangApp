package services

import (
	"auth/config"
	"auth/user/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

func validateUsername(username string) error {
	var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)
	if !usernameRegex.MatchString(username) {
		return errors.New("invalid username: must be 3-20 alphanumeric characters or underscores")
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}

func RegisterUser(username, password string) error {
	if err := validateUsername(username); err != nil {
		return err
	}
	if err := validatePassword(password); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user := models.User{Username: username, Password: string(hashedPassword)}
	result := config.GetDB().Where("username = ?", username).FirstOrCreate(&user)

	if result.RowsAffected == 0 {
		return errors.New("user already exists")
	}

	if result.Error != nil {
		return errors.New("database error: " + result.Error.Error())
	}

	return nil
}

package services

import (
	"auth/config"
	"auth/user/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

// validateUsername ensures the username meets security standards.
func validateUsername(username string) error {
	// Allow only alphanumeric usernames with length between 3 and 20 characters
	var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)
	if !usernameRegex.MatchString(username) {
		return errors.New("invalid username: must be 3-20 alphanumeric characters or underscores")
	}
	return nil
}

// validatePassword ensures password strength
func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	// Additional checks like requiring numbers, special characters, etc., can be added here
	return nil
}

// RegisterUser registers a new user with a hashed password.
func RegisterUser(username, password string) error {
	// Validate inputs
	if err := validateUsername(username); err != nil {
		return err
	}
	if err := validatePassword(password); err != nil {
		return err
	}

	// Hash the password securely
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	// Use FirstOrCreate to handle race conditions
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

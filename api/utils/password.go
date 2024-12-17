package utils

import (
	"errors"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func IsPasswordValid(password string) error {
	// Check length
	if len(password) < 8 || len(password) > 16 {
		return errors.New("password must be between 8 and 16 characters long")
	}

	// Check for at least one digit
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	if !hasNumber {
		return errors.New("password must contain at least one digit")
	}

	// Check for at least one uppercase letter
	hasUpperCase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	if !hasUpperCase {
		return errors.New("password must contain at least one uppercase letter")
	}

	// Check for at least one kind of special character
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()_+{}|:<>?]`).MatchString(password)
	if !hasSpecial {
		return errors.New("password must contain at least one special character")
	}
	return nil
}

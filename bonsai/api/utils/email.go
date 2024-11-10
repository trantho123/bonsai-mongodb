package utils

import (
	"errors"
	"regexp"
	"time"
)

func IsEmailValid(email string) error {
	// Regular expression for validating an email address
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)

	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}

func IsValidDOB(dob string) error {
	// Định dạng mà chúng ta mong muốn
	layout := "2006-01-02" // YYYY-MM-DD

	// Phân tích chuỗi theo định dạng
	_, err := time.Parse(layout, dob)
	if err != nil {
		return errors.New("invalid date format, must be YYYY-MM-DD")
	}

	return nil
}

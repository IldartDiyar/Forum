package service

import (
	"errors"
	"regexp"
	"unicode"
)

var (
	Err_invalid_email    = errors.New("Invalid Email")
	Err_invalid_username = errors.New("Invalid username")
	Err_invalid_password = errors.New("Invalid password")
)

func isValidEmail(email string) error {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{3,29}$`
	re := regexp.MustCompile(pattern)
	if !re.MatchString(email) {
		return Err_invalid_email
	}
	return nil
}

func isValidUsername(username string) error {
	pattern := `^[a-zA-Z][a-zA-Z0-9._-]{3,15}$`
	re := regexp.MustCompile(pattern)
	if !re.MatchString(username) {
		return Err_invalid_username
	}
	return nil
}

func isValidPassword(password string) error {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(password) >= 7 {
		hasMinLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial {
		return nil
	} else {
		return Err_invalid_password
	}

	// re := regexp.MustCompile(`^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$`)

	// if !re.MatchString(password) {
	// 	return Err_invalid_password
	// }
	// return nil
	// hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	// hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	// hasSpecial := regexp.MustCompile(`[@$!%*#?&]`).MatchString(password)

	// if !(hasUppercase && hasDigit && len(password) >= 8) {
	// 	return Err_invalid_password
	// }
}

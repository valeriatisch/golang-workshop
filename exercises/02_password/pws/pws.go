package pws

import "unicode"

func ValidatePassword(password string) bool {

	// Password must be at least 8 characters long
	if len(password) < 8 {
		return false
	}

	// Password must contain at least one uppercase letter, one lowercase letter, and one digit
	var hasUpper, hasLower, hasDigit bool
	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
		} else if unicode.IsLower(char) {
			hasLower = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		}
	}
	if !hasUpper || !hasLower || !hasDigit {
		return false
	}

	// Password is valid
	return true
}
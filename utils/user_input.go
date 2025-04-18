package utils

import (
	"regexp"

	"github.com/dlclark/regexp2"
)

// Regex pour valider le username, l'email et le mot de passe
var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]{3,20}$`)
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func ValidateInputFormat(username string, email string) bool {
	return usernameRegex.MatchString(username) && emailRegex.MatchString(email)
}

func ValidatePassword(password string) bool {
	passwordRegex := regexp2.MustCompile(`^(?=.*[A-Z])(?=.*[0-9])(?=.*[a-z]).{8,}$`, regexp2.None)
	isValid, _ := passwordRegex.MatchString(password)
	return isValid
}

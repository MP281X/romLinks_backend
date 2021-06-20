package encryption

import (
	"golang.org/x/crypto/bcrypt"
)

// hash the password
func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

// compare the hash and the password
func ValidatePassword(password string, hashedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	} else {
		return true
	}
}

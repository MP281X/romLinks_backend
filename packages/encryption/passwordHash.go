package encryption

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// hash the password
func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

// compare the hash and the password
func ValidatePassword(password string, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return errors.New("invalid password")
	}
	return nil
}

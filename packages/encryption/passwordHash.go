package encryption

import (
	"github.com/MP281X/romLinks_backend/packages/logger"
	"golang.org/x/crypto/bcrypt"
)

// hash the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", logger.ErrHashGen
	}
	return string(hashedPassword), nil
}

// compare the hash and the password
func ValidatePassword(password string, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return logger.ErrHashCompare
	}
	return nil
}

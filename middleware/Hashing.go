package middleware

import (
	"golang.org/x/crypto/bcrypt"

)

// HashPassword hashes the password using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePasswords compares the hashed password with the plain text password
func ComparePasswords(userPassword, credsPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(credsPassword))
	if err != nil {
		return false
	}
	return true
}
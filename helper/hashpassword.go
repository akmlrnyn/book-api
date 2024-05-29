package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashPassword), err
}

func VerifyPassword(hashPassword, password string) error{
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
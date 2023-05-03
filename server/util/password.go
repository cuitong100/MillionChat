package util

import "golang.org/x/crypto/bcrypt"
func HashPassword(password string) (string, error) {
	bcrypt.GenerateFormPassword([]byte(password), bcrypt.DefaultCost)
}



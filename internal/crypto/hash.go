package crypto

import "golang.org/x/crypto/bcrypt"

func GenerateHash(value string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
}

func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

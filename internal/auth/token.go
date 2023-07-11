package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gabriel-hahn/devbook/config"
)

func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["userId"] = userID
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString(config.SecretKey)
}

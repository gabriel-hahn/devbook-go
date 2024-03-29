package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gabriel-hahn/devbook/config"
)

func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["userID"] = userID
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString(config.SecretKey)
}

func ValidateToken(r *http.Request) error {
	tokenStr := extractToken(r)
	token, err := jwt.Parse(tokenStr, verificationKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func ExtractUserID(r *http.Request) (uint64, error) {
	tokenStr := extractToken(r)
	token, err := jwt.Parse(tokenStr, verificationKey)
	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIDStr := fmt.Sprintf("%.0f", permissions["userID"])
		userID, err := strconv.ParseUint(userIDStr, 10, 64)
		if err != nil {
			return 0, err
		}

		return userID, nil
	}

	return 0, errors.New("invalid token")
}

func extractToken(r *http.Request) string {
	const validTokenWordCount int = 2

	token := r.Header.Get("Authorization")
	tokenValues := strings.Split(token, " ")

	if len(tokenValues) == validTokenWordCount {
		return tokenValues[1]
	}

	return ""
}

func verificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signature method: %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

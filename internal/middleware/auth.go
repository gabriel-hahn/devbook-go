package middleware

import (
	"log"
	"net/http"

	"github.com/gabriel-hahn/devbook/internal/auth"
	"github.com/gabriel-hahn/devbook/internal/response"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			response.Error(w, http.StatusUnauthorized, err)
			return
		}

		next(w, r)
	}
}

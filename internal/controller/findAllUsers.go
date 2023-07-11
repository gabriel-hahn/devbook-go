package controller

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gabriel-hahn/devbook/internal/database"
	"github.com/gabriel-hahn/devbook/internal/repository"
)

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	if nameOrNick == "" {
		Error(w, http.StatusBadRequest, errors.New("the query param 'user' is required"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	users, err := userRepository.FindAllByFilters(nameOrNick)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusOK, users)
}

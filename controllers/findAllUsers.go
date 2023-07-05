package controllers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gabriel-hahn/devbook/database"
	"github.com/gabriel-hahn/devbook/repositories"
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

	repository := repositories.NewUserRepository(db)
	users, err := repository.FindAllByFilters(nameOrNick)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusOK, users)
}

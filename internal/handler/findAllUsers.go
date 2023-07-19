package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gabriel-hahn/devbook/internal/database"
	"github.com/gabriel-hahn/devbook/internal/repository"
	"github.com/gabriel-hahn/devbook/internal/response"
)

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	if nameOrNick == "" {
		response.Error(w, http.StatusBadRequest, errors.New("the query param 'user' is required"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	users, err := userRepository.FindAllByFilters(nameOrNick)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, users)
}

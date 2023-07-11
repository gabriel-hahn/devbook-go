package controller

import (
	"net/http"
	"strconv"

	"github.com/gabriel-hahn/devbook/internal/database"
	"github.com/gabriel-hahn/devbook/internal/repository"
	"github.com/gorilla/mux"
)

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	if err = userRepository.DeleteByID(userID); err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusNoContent, nil)
}

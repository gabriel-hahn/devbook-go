package controllers

import (
	"net/http"
	"strconv"

	"github.com/gabriel-hahn/devbook/database"
	"github.com/gabriel-hahn/devbook/repositories"
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

	repository := repositories.NewUserRepository(db)
	if err = repository.DeleteByID(userID); err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusNoContent, nil)
}

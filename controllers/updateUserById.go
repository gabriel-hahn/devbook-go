package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gabriel-hahn/devbook/database"
	"github.com/gabriel-hahn/devbook/models"
	"github.com/gabriel-hahn/devbook/repositories"
	"github.com/gorilla/mux"
)

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(models.Update); err != nil {
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
	if err = repository.Update(userID, user); err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusNoContent, nil)
}

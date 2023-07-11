package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gabriel-hahn/devbook/internal/database"
	"github.com/gabriel-hahn/devbook/internal/model"
	"github.com/gabriel-hahn/devbook/internal/repository"
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

	var user model.User
	if err = json.Unmarshal(body, &user); err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(model.Update); err != nil {
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
	if err = userRepository.UpdateByID(userID, user); err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusNoContent, nil)
}

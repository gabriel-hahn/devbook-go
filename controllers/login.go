package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gabriel-hahn/devbook/database"
	"github.com/gabriel-hahn/devbook/internal/crypto"
	"github.com/gabriel-hahn/devbook/models"
	"github.com/gabriel-hahn/devbook/repositories"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	db, err := database.Connect()
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	userFromDB, err := repository.FindByEmail(user.Email)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = crypto.CheckPassword(userFromDB.Password, user.Password); err != nil {
		Error(w, http.StatusUnauthorized, err)
		return
	}
}

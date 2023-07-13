package controller

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gabriel-hahn/devbook/internal/auth"
	"github.com/gabriel-hahn/devbook/internal/crypto"
	"github.com/gabriel-hahn/devbook/internal/database"
	"github.com/gabriel-hahn/devbook/internal/model"
	"github.com/gabriel-hahn/devbook/internal/repository"
	"github.com/gabriel-hahn/devbook/internal/response"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user model.User
	if err = json.Unmarshal(body, &user); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	userFromDB, err := userRepository.FindByEmail(user.Email)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, errors.New("invalid credentials"))
		return
	}

	if err = crypto.CheckPassword(userFromDB.Password, user.Password); err != nil {
		response.Error(w, http.StatusUnauthorized, errors.New("invalid credentials"))
		return
	}

	token, err := auth.CreateToken(userFromDB.ID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(token))
}

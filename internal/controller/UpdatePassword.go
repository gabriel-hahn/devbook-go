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

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var passwordData model.UserPasswordUpdate
	if err := json.Unmarshal(body, &passwordData); err != nil {
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
	dbPassword, err := userRepository.FindPasswordByID(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = crypto.CheckPassword(dbPassword, passwordData.OldPassword); err != nil {
		response.Error(w, http.StatusBadRequest, errors.New("current password is invalid"))
		return
	}

	hashedPassword, err := crypto.GenerateHash(passwordData.NewPassword)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = userRepository.UpdatePassword(userID, string(hashedPassword)); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

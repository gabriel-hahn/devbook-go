package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gabriel-hahn/devbook/database"
	"github.com/gabriel-hahn/devbook/models"
	"github.com/gabriel-hahn/devbook/repositories"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	if err = user.Prepare(models.Signup); err != nil {
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
	user.ID, err = repository.Create(user)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	responseData := models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Nick:      user.Nick,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	JSON(w, http.StatusCreated, responseData)
}

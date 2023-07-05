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

	JSON(w, http.StatusCreated, user)
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find all users"))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get user by ID"))
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user by ID"))
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user by ID"))
}

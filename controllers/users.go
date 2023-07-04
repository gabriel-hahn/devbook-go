package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gabriel-hahn/devbook/database"
	"github.com/gabriel-hahn/devbook/models"
	"github.com/gabriel-hahn/devbook/repositories"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewUserRepository(db)
	repository.Create(user)
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

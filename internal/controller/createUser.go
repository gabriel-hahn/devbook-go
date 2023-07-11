package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gabriel-hahn/devbook/internal/database"
	"github.com/gabriel-hahn/devbook/internal/model"
	"github.com/gabriel-hahn/devbook/internal/repository"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	if err = user.Prepare(model.Signup); err != nil {
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
	user.ID, err = userRepository.Create(user)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	responseData := model.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Nick:  user.Nick,
		Email: user.Email,
	}

	JSON(w, http.StatusCreated, responseData)
}

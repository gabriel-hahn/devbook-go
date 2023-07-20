package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gabriel-hahn/devbook/internal/auth"
	"github.com/gabriel-hahn/devbook/internal/database"
	"github.com/gabriel-hahn/devbook/internal/repository"
	"github.com/gabriel-hahn/devbook/internal/response"
	"github.com/gorilla/mux"
)

func DeletePostById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	userID, err := auth.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepository := repository.NewPostRepository(db)
	postFromDB, err := postRepository.FindByID(postID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if postFromDB.AuthorID != userID {
		response.Error(w, http.StatusForbidden, errors.New("this post does not belong to the current user"))
		return
	}

	if err = postRepository.DeleteByID(postID); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

package handler

import (
	"net/http"
	"strconv"

	"github.com/gabriel-hahn/devbook/internal/database"
	"github.com/gabriel-hahn/devbook/internal/repository"
	"github.com/gabriel-hahn/devbook/internal/response"
	"github.com/gorilla/mux"
)

func LikePostById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepository := repository.NewPostRepository(db)
	if err = postRepository.Like(postID); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

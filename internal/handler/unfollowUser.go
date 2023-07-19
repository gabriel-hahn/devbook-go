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

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)

	unfollowID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if unfollowID == userID {
		response.Error(w, http.StatusForbidden, errors.New("you cannot unfollow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	if err = userRepository.Unfollow(unfollowID, userID); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

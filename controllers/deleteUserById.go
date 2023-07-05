package controllers

import "net/http"

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user by ID"))
}

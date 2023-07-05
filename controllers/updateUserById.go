package controllers

import "net/http"

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user by ID"))
}

package controllers

import "net/http"

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get user by ID"))
}

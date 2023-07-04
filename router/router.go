package router

import "github.com/gorilla/mux"

func Initialize() *mux.Router {
	return mux.NewRouter()
}

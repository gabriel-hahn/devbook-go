package router

import (
	"github.com/gabriel-hahn/devbook/router/routes"
	"github.com/gorilla/mux"
)

func Initialize() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}

package router

import (
	"github.com/gabriel-hahn/devbook/internal/router/route"
	"github.com/gorilla/mux"
)

func Initialize() *mux.Router {
	r := mux.NewRouter()
	return route.Configure(r)
}

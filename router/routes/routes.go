package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI         string
	Method      string
	Callback    func(http.ResponseWriter, *http.Request)
	RequestAuth bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Callback).Methods(route.Method)
	}

	return r
}

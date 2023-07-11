package route

import (
	"net/http"

	"github.com/gabriel-hahn/devbook/internal/middleware"
	"github.com/gorilla/mux"
)

type Route struct {
	URI         string
	Method      string
	Callback    func(http.ResponseWriter, *http.Request)
	RequestAuth bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := append(userRoutes, loginRoutes)

	for _, route := range routes {
		if route.RequestAuth {
			r.HandleFunc(route.URI, middleware.Logger(middleware.Auth(route.Callback))).Methods(route.Method)
			continue
		}

		r.HandleFunc(route.URI, middleware.Logger(route.Callback)).Methods(route.Method)
	}

	return r
}

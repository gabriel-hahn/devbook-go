package routes

import (
	"net/http"

	"github.com/gabriel-hahn/devbook/controllers"
)

var loginRoutes = Route{
	URI:         "/login",
	Method:      http.MethodPost,
	Callback:    controllers.Login,
	RequestAuth: false,
}

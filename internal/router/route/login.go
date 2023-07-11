package route

import (
	"net/http"

	"github.com/gabriel-hahn/devbook/internal/controller"
)

var loginRoutes = Route{
	URI:         "/login",
	Method:      http.MethodPost,
	Callback:    controller.Login,
	RequestAuth: false,
}

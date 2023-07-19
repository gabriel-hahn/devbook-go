package route

import (
	"net/http"

	"github.com/gabriel-hahn/devbook/internal/handler"
)

var loginRoutes = Route{
	URI:         "/login",
	Method:      http.MethodPost,
	Callback:    handler.Login,
	RequestAuth: false,
}

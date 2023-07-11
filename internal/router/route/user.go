package route

import (
	"net/http"

	"github.com/gabriel-hahn/devbook/internal/controller"
)

var userRoutes = []Route{
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Callback:    controller.FindAllUsers,
		RequestAuth: false,
	},
	{
		URI:         "/user",
		Method:      http.MethodPost,
		Callback:    controller.CreateUser,
		RequestAuth: false,
	},
	{
		URI:         "/user/{id}",
		Method:      http.MethodGet,
		Callback:    controller.FindUserById,
		RequestAuth: false,
	},
	{
		URI:         "/user/{id}",
		Method:      http.MethodPut,
		Callback:    controller.UpdateUserById,
		RequestAuth: false,
	},
	{
		URI:         "/user/{id}",
		Method:      http.MethodDelete,
		Callback:    controller.DeleteUserById,
		RequestAuth: false,
	},
}

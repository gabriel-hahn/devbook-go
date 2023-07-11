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
		RequestAuth: true,
	},
	{
		URI:         "/user",
		Method:      http.MethodPost,
		Callback:    controller.CreateUser,
		RequestAuth: true,
	},
	{
		URI:         "/user/{id}",
		Method:      http.MethodGet,
		Callback:    controller.FindUserById,
		RequestAuth: true,
	},
	{
		URI:         "/user/{id}",
		Method:      http.MethodPut,
		Callback:    controller.UpdateUserById,
		RequestAuth: true,
	},
	{
		URI:         "/user/{id}",
		Method:      http.MethodDelete,
		Callback:    controller.DeleteUserById,
		RequestAuth: true,
	},
}

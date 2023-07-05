package routes

import (
	"net/http"

	"github.com/gabriel-hahn/devbook/controllers"
)

var userRoutes = []Route{
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Callback:    controllers.FindAllUsers,
		RequestAuth: false,
	},
	{
		URI:         "/user",
		Method:      http.MethodPost,
		Callback:    controllers.CreateUser,
		RequestAuth: false,
	},
	{
		URI:         "/user/{id}",
		Method:      http.MethodGet,
		Callback:    controllers.FindUserById,
		RequestAuth: false,
	},
	{
		URI:         "/user/{id}",
		Method:      http.MethodPut,
		Callback:    controllers.UpdateUserById,
		RequestAuth: false,
	},
	{
		URI:         "/user/{id}",
		Method:      http.MethodDelete,
		Callback:    controllers.DeleteUserById,
		RequestAuth: false,
	},
}

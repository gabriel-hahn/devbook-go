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
		RequestAuth: false,
	},
	{
		URI:         "/user",
		Method:      http.MethodPut,
		Callback:    controller.UpdateUser,
		RequestAuth: true,
	},
	{
		URI:         "/user",
		Method:      http.MethodDelete,
		Callback:    controller.DeleteUser,
		RequestAuth: true,
	},
	{
		URI:         "/user/{id}",
		Method:      http.MethodGet,
		Callback:    controller.FindUserById,
		RequestAuth: true,
	},
	{
		URI:         "/user/{userId}/follow",
		Method:      http.MethodPost,
		Callback:    controller.FollowUser,
		RequestAuth: true,
	},
	{
		URI:         "/user/{userId}/unfollow",
		Method:      http.MethodPost,
		Callback:    controller.UnfollowUser,
		RequestAuth: true,
	},
	{
		URI:         "/user/{userId}/followers",
		Method:      http.MethodGet,
		Callback:    controller.FindAllFollowers,
		RequestAuth: true,
	},
	{
		URI:         "/user/{userId}/following",
		Method:      http.MethodGet,
		Callback:    controller.FindAllFollowing,
		RequestAuth: true,
	},
	{
		URI:         "/user/update-password",
		Method:      http.MethodPost,
		Callback:    controller.UpdatePassword,
		RequestAuth: true,
	},
}

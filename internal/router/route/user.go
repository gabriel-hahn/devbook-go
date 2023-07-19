package route

import (
	"net/http"

	"github.com/gabriel-hahn/devbook/internal/handler"
)

var userRoutes = []Route{
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Callback:    handler.FindAllUsers,
		RequestAuth: true,
	},
	{
		URI:         "/user",
		Method:      http.MethodPost,
		Callback:    handler.CreateUser,
		RequestAuth: false,
	},
	{
		URI:         "/user",
		Method:      http.MethodPut,
		Callback:    handler.UpdateUser,
		RequestAuth: true,
	},
	{
		URI:         "/user",
		Method:      http.MethodDelete,
		Callback:    handler.DeleteUser,
		RequestAuth: true,
	},
	{
		URI:         "/user/{id}",
		Method:      http.MethodGet,
		Callback:    handler.FindUserById,
		RequestAuth: true,
	},
	{
		URI:         "/user/{userId}/follow",
		Method:      http.MethodPost,
		Callback:    handler.FollowUser,
		RequestAuth: true,
	},
	{
		URI:         "/user/{userId}/unfollow",
		Method:      http.MethodPost,
		Callback:    handler.UnfollowUser,
		RequestAuth: true,
	},
	{
		URI:         "/user/{userId}/followers",
		Method:      http.MethodGet,
		Callback:    handler.FindAllFollowers,
		RequestAuth: true,
	},
	{
		URI:         "/user/{userId}/following",
		Method:      http.MethodGet,
		Callback:    handler.FindAllFollowing,
		RequestAuth: true,
	},
	{
		URI:         "/user/update-password",
		Method:      http.MethodPost,
		Callback:    handler.UpdatePassword,
		RequestAuth: true,
	},
}

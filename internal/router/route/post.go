package route

import (
	"net/http"

	"github.com/gabriel-hahn/devbook/internal/controller"
)

var postRoutes = []Route{
	{
		URI:         "/posts",
		Method:      http.MethodGet,
		Callback:    controller.FindAllPosts,
		RequestAuth: true,
	},
	{
		URI:         "/post/{postId}",
		Method:      http.MethodGet,
		Callback:    controller.FindPostById,
		RequestAuth: true,
	},
	{
		URI:         "/post/{postId}",
		Method:      http.MethodPut,
		Callback:    controller.UpdatePostById,
		RequestAuth: true,
	},
	{
		URI:         "/post/{postId}",
		Method:      http.MethodDelete,
		Callback:    controller.DeletePostById,
		RequestAuth: true,
	},
	{
		URI:         "/post",
		Method:      http.MethodPost,
		Callback:    controller.CreatePost,
		RequestAuth: true,
	},
}

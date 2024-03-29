package route

import (
	"net/http"

	"github.com/gabriel-hahn/devbook/internal/handler"
)

var postRoutes = []Route{
	{
		URI:         "/posts",
		Method:      http.MethodGet,
		Callback:    handler.FindAllPosts,
		RequestAuth: true,
	},
	{
		URI:         "/users/{userId}/posts",
		Method:      http.MethodGet,
		Callback:    handler.FindAllPostsByUser,
		RequestAuth: true,
	},
	{
		URI:         "/post/{postId}",
		Method:      http.MethodGet,
		Callback:    handler.FindPostById,
		RequestAuth: true,
	},
	{
		URI:         "/post/{postId}",
		Method:      http.MethodPut,
		Callback:    handler.UpdatePostById,
		RequestAuth: true,
	},
	{
		URI:         "/post/{postId}",
		Method:      http.MethodDelete,
		Callback:    handler.DeletePostById,
		RequestAuth: true,
	},
	{
		URI:         "/post",
		Method:      http.MethodPost,
		Callback:    handler.CreatePost,
		RequestAuth: true,
	},
	{
		URI:         "/post/{postId}/like",
		Method:      http.MethodPost,
		Callback:    handler.LikePostById,
		RequestAuth: true,
	},
	{
		URI:         "/post/{postId}/dislike",
		Method:      http.MethodPost,
		Callback:    handler.DislikePostById,
		RequestAuth: true,
	},
}

package routes

import (
	"devbook_api/src/controllers"
	"net/http"
)

var postRoutes = []Route{
	{
		URI:          "/api/posts",
		Method:       http.MethodPost,
		Function:     controllers.CreatePost,
		RequiresAuth: true,
	},
	{
		URI:          "/api/posts",
		Method:       http.MethodGet,
		Function:     controllers.FindPosts,
		RequiresAuth: false,
	},
	{
		URI:          "/api/posts/{postID}",
		Method:       http.MethodGet,
		Function:     controllers.FindPost,
		RequiresAuth: true,
	},
	{
		URI:          "/api/posts/{postID}",
		Method:       http.MethodPut,
		Function:     controllers.UpdatePost,
		RequiresAuth: true,
	},
	{
		URI:          "/api/posts/{postID}",
		Method:       http.MethodDelete,
		Function:     controllers.DeletePost,
		RequiresAuth: true,
	},
	{
		URI:          "/api/user/{userID}/posts",
		Method:       http.MethodGet,
		Function:     controllers.FindPostsByUser,
		RequiresAuth: false,
	},
	{
		URI:          "/api/posts/{postID}/like",
		Method:       http.MethodPost,
		Function:     controllers.LikePost,
		RequiresAuth: true,
	},
	{
		URI:          "/api/posts/{postID}/like",
		Method:       http.MethodDelete,
		Function:     controllers.UnlikePost,
		RequiresAuth: true,
	},
}

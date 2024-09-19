package routes

import (
	"devbook_app/src/controllers"
	"net/http"
)

var postRoutes = []Route{
	// {
	// 	URI:          "/posts",
	// 	Method:       http.MethodGet,
	// 	Function:     controllers.GetPosts,
	// 	RequiresAuth: true,
	// },
	{
		URI:          "/posts",
		Method:       http.MethodPost,
		Function:     controllers.CreatePost,
		RequiresAuth: true,
	},
	// {
	// 	URI:          "/posts/{postID}",
	// 	Method:       http.MethodGet,
	// 	Function:     controllers.GetPost,
	// 	RequiresAuth: true,
	// },
	{
		URI:          "/posts/{postID}",
		Method:       http.MethodPut,
		Function:     controllers.UpdatePost,
		RequiresAuth: true,
	},
	{
		URI:          "/posts/{postID}",
		Method:       http.MethodDelete,
		Function:     controllers.DeletePost,
		RequiresAuth: true,
	},
	{
		URI:          "/posts/{postID}/like",
		Method:       http.MethodPost,
		Function:     controllers.LikePost,
		RequiresAuth: true,
	},
	{
		URI:          "/posts/{postID}/like",
		Method:       http.MethodDelete,
		Function:     controllers.UnlikePost,
		RequiresAuth: true,
	},
	{
		URI:          "/create",
		Method:       http.MethodGet,
		Function:     controllers.LoadCreatePostView,
		RequiresAuth: true,
	},
	{
		URI:          "/edit/{postID}",
		Method:       http.MethodGet,
		Function:     controllers.LoadEditPostView,
		RequiresAuth: true,
	},
}

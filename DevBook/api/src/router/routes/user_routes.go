package routes

import (
	"devbook_api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:          "/api/user",
		Method:       http.MethodGet,
		Function:     controllers.GetUsers,
		RequiresAuth: true,
	},
	{
		URI:          "/api/user/{id}",
		Method:       http.MethodGet,
		Function:     controllers.GetUser,
		RequiresAuth: false,
	},
	{
		URI:          "/api/user",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		RequiresAuth: false,
	},
	{
		URI:          "/api/user/{id}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		RequiresAuth: false,
	},
	{
		URI:          "/api/user/{id}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		RequiresAuth: false,
	},
	{
		URI:          "/api/user/{id}/follow",
		Method:       http.MethodPost,
		Function:     controllers.FollowUser,
		RequiresAuth: true,
	},
	{
		URI:          "/api/user/{id}/unfollow",
		Method:       http.MethodPost,
		Function:     controllers.UnfollowUser,
		RequiresAuth: true,
	},
	{
		URI:          "/api/user/{id}/followers",
		Method:       http.MethodGet,
		Function:     controllers.GetFollowers,
		RequiresAuth: true,
	},
	{
		URI:          "/api/user/{id}/following",
		Method:       http.MethodGet,
		Function:     controllers.GetFollowing,
		RequiresAuth: true,
	},
	{
		URI:          "/api/user/{id}/updatePassword",
		Method:       http.MethodPost,
		Function:     controllers.UpdatePassword,
		RequiresAuth: true,
	},
}

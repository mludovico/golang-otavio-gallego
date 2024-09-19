package routes

import (
	"devbook_app/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:          "/signup",
		Method:       http.MethodGet,
		Function:     controllers.Signup,
		RequiresAuth: false,
	},
	{
		URI:          "/signup",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		RequiresAuth: false,
	},
	{
		URI:          "/search",
		Method:       http.MethodGet,
		Function:     controllers.SearchUsers,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userID}",
		Method:       http.MethodGet,
		Function:     controllers.GetUser,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userID}/follow",
		Method:       http.MethodPost,
		Function:     controllers.FollowUser,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userID}/unfollow",
		Method:       http.MethodPost,
		Function:     controllers.UnfollowUser,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userID}/followers",
		Method:       http.MethodGet,
		Function:     controllers.ListFollowers,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userID}/following",
		Method:       http.MethodGet,
		Function:     controllers.ListFollowing,
		RequiresAuth: true,
	},
	{
		URI:          "/profile",
		Method:       http.MethodGet,
		Function:     controllers.Profile,
		RequiresAuth: true,
	},
	{
		URI:          "/profile/update",
		Method:       http.MethodGet,
		Function:     controllers.EditProfile,
		RequiresAuth: true,
	},
	{
		URI:          "/profile/update",
		Method:       http.MethodPost,
		Function:     controllers.UpdateProfile,
		RequiresAuth: true,
	},
	{
		URI:          "/profile/password",
		Method:       http.MethodGet,
		Function:     controllers.EditPassword,
		RequiresAuth: true,
	},
	{
		URI:          "/profile/password",
		Method:       http.MethodPost,
		Function:     controllers.UpdatePassword,
		RequiresAuth: true,
	},
	{
		URI:          "/profile/delete",
		Method:       http.MethodPost,
		Function:     controllers.DeleteAccount,
		RequiresAuth: true,
	},
}

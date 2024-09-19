package routes

import (
	"devbook_app/src/controllers"
	"net/http"
)

var loginRoutes = []Route{
	{
		URI:          "/login",
		Method:       http.MethodGet,
		Function:     controllers.Login,
		RequiresAuth: false,
	},
	{
		URI:          "/login",
		Method:       http.MethodPost,
		Function:     controllers.Authenticate,
		RequiresAuth: false,
	},
	{
		URI:          "/logout",
		Method:       http.MethodGet,
		Function:     controllers.Logout,
		RequiresAuth: false,
	}}

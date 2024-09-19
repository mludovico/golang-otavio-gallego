package routes

import (
	"devbook_app/src/controllers"
	"net/http"
)

var homeRoutes = []Route{
	{
		URI:          "/",
		Method:       http.MethodGet,
		Function:     controllers.Home,
		RequiresAuth: true,
	},
	{
		URI:          "/home",
		Method:       http.MethodGet,
		Function:     controllers.Home,
		RequiresAuth: true,
	},
}

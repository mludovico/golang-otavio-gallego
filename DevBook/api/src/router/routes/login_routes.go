package routes

import (
	"devbook_api/src/controllers"
	"net/http"
)

var loginRoutes = []Route{
	{
		URI:          "/api/login",
		Method:       http.MethodPost,
		Function:     controllers.Login,
		RequiresAuth: false,
	},
}

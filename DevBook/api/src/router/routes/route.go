package routes

import (
	"devbook_api/src/middlleware"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI          string
	Method       string
	Function     func(w http.ResponseWriter, r *http.Request)
	RequiresAuth bool
}

func Configure(router *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoutes...)
	routes = append(routes, postRoutes...)
	for _, route := range routes {
		if route.RequiresAuth {
			router.HandleFunc(route.URI, middlleware.Logger(
				middlleware.Cors(
					middlleware.Authenticate(route.Function),
				),
			)).Methods(route.Method, http.MethodOptions)
			continue
		} else {
			router.HandleFunc(route.URI, middlleware.Logger(
				middlleware.Cors(
					route.Function,
				),
			)).Methods(route.Method, http.MethodOptions)
		}
	}
	return router
}

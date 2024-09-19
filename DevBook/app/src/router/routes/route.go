package routes

import (
	"devbook_app/src/middlewares"
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
	routes := loginRoutes
	routes = append(routes, userRoutes...)
	routes = append(routes, homeRoutes...)
	routes = append(routes, postRoutes...)
	for _, route := range routes {
		if route.RequiresAuth {
			router.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
			continue
		} else {
			router.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}

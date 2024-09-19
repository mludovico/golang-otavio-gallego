package router

import (
	"devbook_app/src/router/routes"

	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	return routes.Configure(router)
}

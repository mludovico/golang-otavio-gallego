package router

import (
	"devbook_api/src/router/routes"

	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	return routes.Configure(router)
}

package main

import (
	"devbook_api/src/config"
	"devbook_api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()
	router := router.GetRouter()
	fmt.Printf("Server running on port %d\n", config.ApiPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), router))
}

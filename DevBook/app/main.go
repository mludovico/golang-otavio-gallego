package main

import (
	"devbook_app/src/config"
	"devbook_app/src/cookies"
	"devbook_app/src/router"
	"devbook_app/src/utils"
	"fmt"
	"log"
	"net/http"
)

const port = 3000

func main() {
	config.LoadConfig()
	cookies.Configure()
	fmt.Printf("Config loaded: ApiURL=%s, ApiPort=%d, HashKey=%s, BlockKey=%s\n", config.APIURL, config.Port, config.HashKey, config.BlockKey)
	utils.LoadTemplates()

	router := router.GetRouter()

	fmt.Printf("App running on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
